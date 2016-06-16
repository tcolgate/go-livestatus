package livestatus

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"time"
)

var (
	acknowledge_host_problem = template.Must(template.New("ACKNOWLEDGE_HOST_PROBLEM").Parse("{{.HostName}};{{.Sticky}};{{.Notify}};{{.Persistent}};{{.Author}};{{.Comment}}"))
	acknowledge_svc_problem  = template.Must(template.New("ACKNOWLEDGE_SVC_PROBLEM").Parse("{{.HostName}};{{.ServiceDescription}};{{.Sticky}};{{.Notify}};{{.Persistent}};{{.Author}};{{.Comment}}"))
	schedule_host_downtime   = template.Must(template.New("SCHEDULE_HOST_DOWNTIME").Parse("{{.HostName}};{{.StartTime}};{{.EndTime}};{{.Fixed}};{{.TriggerID}};{{.Duration}};{{.Author}};{{.Comment}}"))
	schedule_svc_downtime    = template.Must(template.New("SCHEDULE_SVC_DOWNTIME").Parse("{{.HostName}};{{.ServiceDesription}};{{.StartTime}};{{.EndTime}};{{.Fixed}};{{.TriggerID}};{{.Duration}};{{.Author}};{{.Comment}}"))
)

// Command is a binding command instance.
type Command struct {
	tmpl *template.Template
	args *commandArgs
	ls   *Livestatus
}

type stickyBool bool

func (b stickyBool) String() string {
	if b {
		return "2"
	} else {
		return "0"
	}
}

type persistentBool bool

func (b persistentBool) String() string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

type notifyBool bool

func (b notifyBool) String() string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

type stringDate struct{ time.Time }

func (d stringDate) String() string {
	return fmt.Sprintf("%d", d.Unix())
}

type stringDuration struct{ time.Duration }

func (d stringDuration) String() string {
	return fmt.Sprintf("%d", d.Duration/time.Second)
}

type commandArgs struct {
	HostName           string
	ServiceDescription string
	Sticky             stickyBool
	Notify             notifyBool
	Persistent         persistentBool
	Author             string
	Comment            string
	StartTime          int
	EndTime            int
	Duration           int
	TriggerID          int
	DowntimeID         int
}

// Exec executes the query.
func (c *Command) Exec() (*Response, error) {
	resp := &Response{}

	// Connect to socket
	conn, err := c.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Send command data
	cmd, err := c.buildCmd(time.Now())
	if err != nil {
		return nil, err
	}

	conn.Write([]byte(cmd))

	// Read response header
	data := make([]byte, 16)

	_, err = conn.Read(data)
	if err != nil {
		return nil, err
	}

	log.Println(string(data))
	return resp, nil
}

func (c *Command) buildCmd(t time.Time) (string, error) {
	buf := &bytes.Buffer{}

	if c.tmpl == nil {
	}
	if err := c.tmpl.Execute(buf, c.args); err != nil {
		return "", err
	}

	return fmt.Sprintf("COMMAND %d %s;%s\n\n", t.Unix(), c.tmpl.Name(), buf.String()), nil
}

func (c *Command) dial() (net.Conn, error) {
	if c.ls.dialer != nil {
		return c.ls.dialer()
	} else {
		return net.Dial(c.ls.network, c.ls.address)
	}
}

func newCommand(ls *Livestatus) *Command {
	return &Command{
		nil,
		nil,
		ls,
	}
}

func (c *Command) AcknowledgeHostProblem(name string, sticky, notify, persistent bool, author, comment string) *Command {
	return &Command{
		tmpl: acknowledge_host_problem,
		args: &commandArgs{
			HostName:   name,
			Sticky:     stickyBool(sticky),
			Notify:     notifyBool(notify),
			Persistent: persistentBool(persistent),
			Author:     author,
			Comment:    comment,
		},
		ls: c.ls,
	}
}
