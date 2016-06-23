package livestatus

import (
	"fmt"
	"net"
	"strings"
	"time"
)

/*
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
*/

// Command is a binding command instance.
type Command struct {
	cmd  string
	vals []string
	ls   *Livestatus
}

func newCommand(ls *Livestatus) *Command {
	return &Command{
		ls: ls,
	}
}

func (c *Command) Raw(cmd string) {
	c.cmd = cmd
}

func (c *Command) Arg(v interface{}) {
	c.vals = append(c.vals, fmt.Sprintf("%v", v))
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
	c.Arg(s)
}

func (c *Command) ServiceDescription(s string) {
	c.Arg(s)
}

func (c *Command) Sticky(b bool) {
	c.Arg(stickyBool(b).String())
}

func (c *Command) Notify(b bool) {
	c.Arg(normalBool(b).String())
}

func (c *Command) Fixed(b bool) {
	c.Arg(normalBool(b).String())
}

func (c *Command) Persistent(b bool) {
	c.Arg(normalBool(b).String())
}

func (c *Command) Author(s string) {
	c.Arg(s)
}

func (c *Command) Comment(s string) {
	c.Arg(s)
}

func (c *Command) Start(t time.Time) {
	c.Arg(t.Unix())
}

func (c *Command) End(t time.Time) {
	c.Arg(t.Unix())
}

func (c *Command) Duration(t time.Duration) {
	c.Arg(stringDuration{t}.String())
}

func (c *Command) TriggerID(i int) {
	c.Arg(i)
}

func (c *Command) DowntimeID(i int) {
	c.Arg(i)
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
	cmdStr := fmt.Sprintf("COMMAND [%d] %s", t.Unix(), c.cmd)
	cmdStr = fmt.Sprintf("%s;%s", cmdStr, strings.Join(c.vals, ";"))

	return fmt.Sprintf("%s\n", cmdStr), nil
}

func (c *Command) dial() (net.Conn, error) {
	if c.ls.dialer != nil {
		return c.ls.dialer()
	} else {
		return net.Dial(c.ls.network, c.ls.address)
	}
}
