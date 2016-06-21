package livestatus

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

var ()

func init() {
	// These service definition are taken from https://old.nagios.org/developerinfo/externalcommands
	addCommands(
		"ACKNOWLEDGE_HOST_PROBLEM;<host_name>;<sticky>;<notify>;<persistent>;<author>;<comment>",
		"ACKNOWLEDGE_SVC_PROBLEM;<host_name>;<service_description>;<sticky>;<notify>;<persistent>;<author>;<comment>",
		"SCHEDULE_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>",
		"SCHEDULE_SVC_DOWNTIME;<host_name>;<service_desription>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>",
		"SCHEDULE_HOST_SVC_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>")
}

type commandDef struct {
	name  string
	args  []string
	needs map[string]bool
}

var commands map[string]*commandDef

var cmdTmpl = regexp.MustCompile("^[A-Z_]+$")
var argTmpl = regexp.MustCompile("^<([a-z_]+)>$")

func addCommands(ss ...string) {
	for _, s := range ss {
		if commands == nil {
			commands = make(map[string]*commandDef)
		}
		def := mustParseCommandDef(s)
		commands[def.name] = def
	}
}

// mustParseCommandDef parses a command definition from nagios. Panics if it cannot
func mustParseCommandDef(s string) *commandDef {
	parts := strings.Split(s, ";")
	if !cmdTmpl.MatchString(parts[0]) {
		panic("malformed command definition")
	}
	cmd := parts[0]

	args := []string{}
	for _, a := range parts[1:] {
		ms := argTmpl.FindStringSubmatch(a)
		if ms == nil {
			panic("malformed command definition")
		}
		args = append(args, ms[1])
	}

	c := commandDef{
		name:  cmd,
		args:  args,
		needs: map[string]bool{},
	}
	for _, a := range c.args {
		c.needs[a] = true
	}
	return &c
}

// Command is a binding command instance.
type Command struct {
	cmd  string
	vals map[string]string
	ls   *Livestatus
}

func newCommand(s string, ls *Livestatus) *Command {
	return &Command{
		cmd:  s,
		vals: map[string]string{},
		ls:   ls,
	}
}

func (c *Command) setVal(k string, v interface{}) {
	c.vals[k] = fmt.Sprintf("%v", v)
}

type stickyBool bool

func (b stickyBool) String() string {
	if b {
		return "2"
	} else {
		return "0"
	}
}

type normalBool bool

func (b normalBool) String() string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

type stringDuration struct{ time.Duration }

func (d stringDuration) String() string {
	return fmt.Sprintf("%d", d.Duration/time.Second)
}

// KeepAliveOff disables the default keepalive from Command
func (q *Query) KeepAliveOff() *Query {
	q.ls.keepalive = false
	return q
}

func (c *Command) Hostname(s string) {
	c.setVal("host_name", s)
}

func (c *Command) ServiceDescription(s string) {
	c.setVal("service_description", s)
}

func (c *Command) Sticky(b bool) {
	c.setVal("sticky", stickyBool(b).String())
}

func (c *Command) Notify(b bool) {
	c.setVal("notify", normalBool(b).String())
}

func (c *Command) Fixed(b bool) {
	c.setVal("fixed", normalBool(b).String())
}

func (c *Command) Persistent(b bool) {
	c.setVal("persistent", normalBool(b).String())
}

func (c *Command) Author(s string) {
	c.setVal("author", s)
}

func (c *Command) Comment(s string) {
	c.setVal("comment", s)
}

func (c *Command) Start(t time.Time) {
	c.setVal("start_time", t.Unix())
}

func (c *Command) End(t time.Time) {
	c.setVal("end_time", t.Unix())
}

func (c *Command) Duration(t time.Duration) {
	c.setVal("duration", stringDuration{t}.String())
}

func (c *Command) TriggerID(i int) {
	c.setVal("trigger_id", i)
}

func (c *Command) DowntimeID(i int) {
	c.setVal("downtime_id", i)
}

// Exec executes the query.
func (c *Command) Exec() (*Response, error) {
	resp := &Response{}

	var err error
	var conn net.Conn

	if c.ls.keepConn != nil {
		conn = c.ls.keepConn
	} else {
		// Connect to socket
		conn, err = c.dial()
		if err != nil {
			return nil, err
		}
	}

	if !c.ls.keepalive {
		c.ls.keepConn = nil
		defer conn.Close()
	} else {
		c.ls.keepConn = conn
	}

	// Send command data
	cmd, err := c.buildCmd(time.Now())
	if err != nil {
		return nil, err
	}

	conn.Write([]byte(cmd))
	// You get nothing back from an external command
	// no way of knowing if this has worked

	return resp, nil
}

func (c *Command) buildCmd(t time.Time) (string, error) {
	//verify the command has all the args set, report what is missing
	def, ok := commands[c.cmd]
	if !ok {
		// probably could fall back to dumps the vals out in the order provided
		return "", fmt.Errorf("unknown command %s", c.cmd)
	}

	provided := map[string]bool{}
	// Check the args are needed
	for a := range c.vals {
		if n, ok := def.needs[a]; !ok || !n {
			return "", fmt.Errorf("command %s does not need argument %s", c.cmd, a)
		}
		provided[a] = true
	}

	// Check the needed args are provided
	for n := range def.needs {
		if p, ok := provided[n]; !ok || !p {
			return "", fmt.Errorf("command %s requires argument %s", c.cmd, n)
		}
	}

	cmdStr := fmt.Sprintf("COMMAND [%d] %s", t.Unix(), c.cmd)
	for _, a := range def.args {
		cmdStr = fmt.Sprintf("%s;%s", cmdStr, c.vals[a])
	}

	return fmt.Sprintf("%s\n", cmdStr), nil
}

func (c *Command) dial() (net.Conn, error) {
	if c.ls.dialer != nil {
		return c.ls.dialer()
	} else {
		return net.Dial(c.ls.network, c.ls.address)
	}
}
