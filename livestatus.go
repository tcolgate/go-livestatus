// Package livestatus provides a binding to MK Livestatus sockets.
package livestatus

import "net"

// Livestatus is a binding instance.
type Livestatus struct {
	network string
	address string
	dialer  func() (net.Conn, error)
}

// Query creates a new query instance on a spacific table.
func (l *Livestatus) Query(table string) *Query {
	return newQuery(table, l)
}

// Command creates a new command instanc.
func (l *Livestatus) Command(cmd string) *Command {
	return newCommand(cmd, l)
}

// NewLivestatus creates a new binding instance.
func NewLivestatus(network, address string) *Livestatus {
	return &Livestatus{
		network: network,
		address: address,
	}
}

// NewLivestatusWithDialer creates a new binding that uses the net.Conn returned
// by the provided dialer function.
func NewLivestatusWithDialer(dialer func() (net.Conn, error)) *Livestatus {
	return &Livestatus{
		dialer: dialer,
	}
}
