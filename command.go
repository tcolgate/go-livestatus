package livestatus

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	commandCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "livestatus_command_count",
		Help: "",
	}, []string{"command"})
)

func init() {
	prometheus.MustRegister(commandCount)
}

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

type CommandOpFunc func(*Command)

func (c *Command) Op(op CommandOpFunc) {
	op(c)
}

// KeepAliveOff disables the default keepalive from Command
func (q *Query) KeepAliveOff() *Query {
	q.ls.keepalive = false
	return q
}

// Exec executes the query.
func (c *Command) Exec() (*Response, error) {
	resp := &Response{}

	var err error
	var conn net.Conn

	if c.ls.keepConn != nil {
		conn = c.ls.keepConn
		connectReuseCount.
			WithLabelValues(conn.RemoteAddr().String()).
			Inc()
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

	commandCount.WithLabelValues(c.cmd).Inc()

	return resp, nil
}

func (c *Command) buildCmd(t time.Time) (string, error) {
	cmdStr := fmt.Sprintf("COMMAND [%d] %s", t.Unix(), c.cmd)
	cmdStr = fmt.Sprintf("%s;%s", cmdStr, strings.Join(c.vals, ";"))

	return fmt.Sprintf("%s\n", cmdStr), nil
}

func (c *Command) dial() (cc net.Conn, err error) {
	defer func() {
		if err != nil {
			connectCount.
				WithLabelValues(cc.RemoteAddr().String()).
				Inc()
			return
		}
		connectErrCount.
			WithLabelValues(err.Error()).
			Inc()
	}()

	if c.ls.dialer != nil {
		return c.ls.dialer()
	} else {
		return net.Dial(c.ls.network, c.ls.address)
	}
}
