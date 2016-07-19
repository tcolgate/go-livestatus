package livestatus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	queryDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "livestatus_query_duration_seconds",
		Help:    "Histogram of successful livestatus query durations, waiting indicates a WaitCondition in use",
		Buckets: prometheus.LinearBuckets(0, 0.2, 10),
	}, []string{"table", "waiting"})

	responseSize = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "livestatus_response_size_bytes",
		Help:    "Histogram of successful livestatus query response sizes",
		Buckets: prometheus.ExponentialBuckets(32, 4, 8),
	}, []string{"table"})

	queryCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "livestatus_query_count",
		Help: "Count of the livestatus queries sent",
	}, []string{"table", "status", "error"})

	connectCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "livestatus_connect_count",
		Help: "Count of the successful connections",
	}, []string{"addr"})

	connectReuseCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "livestatus_connect_reuse_count",
		Help: "Count of the number of times a connection is reused",
	}, []string{"addr"})

	connectErrCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "livestatus_connect_error_count",
		Help: "Count of the failed connection attempts",
	}, []string{"error"})
)

func init() {
	prometheus.MustRegister(queryDuration)
	prometheus.MustRegister(queryCount)
	prometheus.MustRegister(responseSize)
	prometheus.MustRegister(connectCount)
	prometheus.MustRegister(connectReuseCount)
	prometheus.MustRegister(connectErrCount)
}

// Query is a binding query instance.
type Query struct {
	table   string
	headers []string
	columns []string
	ls      *Livestatus
	waiting bool
}

// Columns sets the names of the columns to retrieve when executing a query.
func (q *Query) Columns(names ...string) *Query {
	q.headers = append(q.headers, "Columns: "+strings.Join(names, " "))
	q.columns = names
	return q
}

// KeepAlive keeps the connection open after the query, for re-use
func (q *Query) KeepAlive() *Query {
	q.headers = append(q.headers, "KeepAlive: on")
	return q
}

// Filter sets a new filter to apply to the query.
func (q *Query) Filter(rule string) *Query {
	q.headers = append(q.headers, "Filter: "+rule)
	return q
}

// And combines the n last filters into a new filter using a `And` operation.
func (q *Query) And(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("And: %d", n))
	return q
}

// Or combines the n last filters into a new filter using a `Or` operation.
func (q *Query) Or(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("Or: %d", n))
	return q
}

// Negate negates the most recent filter.
func (q *Query) Negate() *Query {
	q.headers = append(q.headers, "Negate:")
	return q
}

// Limit the query to n responses.
func (q *Query) Limit(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("Limit: %d", n))
	return q
}

// WaitObject sets the object within the queried table to wait on. For the table
// hosts, hostgroups, servicegroups, contacts and contactgroups this is simply
// the name of the object. For the table services it is the hostname followed
// by a space followed by the service description
func (q *Query) WaitObject(obj string) *Query {
	q.headers = append(q.headers, "WaitObject: "+obj)
	return q
}

// WaitCondition sets a new wait condition  to apply to the query.
func (q *Query) WaitCondition(rule string) *Query {
	q.waiting = true
	q.headers = append(q.headers, "WaitCondition: "+rule)
	return q
}

// WaitConditionAnd combines the n last wait conditions into a new wait
// condition using a `And` operation.
func (q *Query) WaitConditionAnd(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitConditionAnd: %d", n))
	return q
}

// WaitConditionOr combines the n last wait condition into a new wait condition
// using a `Or` operation.
func (q *Query) WaitConditionOr(n int) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitConditionOr: %d", n))
	return q
}

// WaitConditionNegate negates the most recent wait condition.
func (q *Query) WaitConditionNegate() *Query {
	q.headers = append(q.headers, "WaitConditionNegate:")
	return q
}

// WaitTrigger sets the nagios event that will trigger a check of
// the wait condition.
func (q *Query) WaitTrigger(event string) *Query {
	q.headers = append(q.headers, "WaitTrigger: "+event)
	return q
}

// WaitTimeout set a timeout for the wait condition.
func (q *Query) WaitTimeout(t time.Duration) *Query {
	q.headers = append(q.headers, fmt.Sprintf("WaitTimeout: %d", t/time.Millisecond))
	return q
}

// Exec executes the query.
func (q *Query) Exec() (*Response, error) {
	var err error
	var conn net.Conn

	resp := &Response{}
	st := time.Now()
	size := 0

	buf := bytes.NewBuffer(nil)

	defer func() {
		errstring := "success"
		if err != nil {
			errstring = err.Error()
		}
		queryCount.WithLabelValues(q.table, fmt.Sprintf("%d", resp.Status), errstring).Inc()

		if err == nil {
			queryDuration.
				WithLabelValues(q.table, fmt.Sprintf("%t", q.waiting)).
				Observe(float64(time.Now().Sub(st)) / float64(time.Second))
			responseSize.
				WithLabelValues(q.table).
				Observe(float64(size))
		}
	}()

	if q.ls.keepConn != nil {
		conn = q.ls.keepConn
		connectReuseCount.
			WithLabelValues(conn.RemoteAddr().String()).
			Inc()
	} else {
		// Connect to socket
		conn, err = q.dial()
		if err != nil {
			return nil, err
		}
	}

	if !q.ls.keepalive {
		q.ls.keepConn = nil
		defer conn.Close()
	} else {
		q.ls.keepConn = conn
	}

	// Send command data
	conn.Write([]byte(q.buildCmd()))

	data := make([]byte, 16)
	_, err = conn.Read(data)
	if err != nil {
		return nil, err
	}

	resp.Status, err = strconv.Atoi(string(data[:3]))
	if err != nil {
		return nil, err
	}

	for {
		data = make([]byte, 1024)

		n, err := conn.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		buf.Write(bytes.TrimRight(data, "\x00"))

		// New line signals the end of content. This check helps
		// if the connection is not forcibly closed
		if data[n-1] == byte('\n') {
			break
		}
	}

	if buf.Len() == 0 {
		return resp, nil
	}
	size = buf.Len()

	// Parse received data for records
	resp.Records, err = q.parse(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *Query) buildCmd() string {
	cmd := "GET " + q.table

	// Append headers if any
	if len(q.headers) > 0 {
		cmd += "\n" + strings.Join(q.headers, "\n")
	}

	// Set default headers
	cmd += "\nResponseHeader: fixed16"
	cmd += "\nOutputFormat: json"
	cmd += "\n\n"

	return cmd
}

func (q *Query) dial() (c net.Conn, err error) {
	defer func() {
		if err == nil {
			connectCount.
				WithLabelValues(c.RemoteAddr().String()).
				Inc()
			return
		}
		connectErrCount.
			WithLabelValues(err.Error()).
			Inc()
	}()

	if q.ls.dialer != nil {
		return q.ls.dialer()
	} else {
		return net.Dial(q.ls.network, q.ls.address)
	}
}

func (q *Query) parse(data []byte) ([]Record, error) {
	var (
		records []Record
		rows    [][]interface{}
	)

	// Unmarshal received data
	if err := json.Unmarshal(data, &rows); err != nil {
		str := string(data)
		sz := len(data)
		if sz > 128 {
			str = string(data[0:127]) + "..."
		}
		return nil, errors.New(str)
	} else if len(q.columns) == 0 && len(rows) < 2 || len(q.columns) > 0 && len(rows) < 1 {
		return records, nil
	}

	// Extract columns names from first row
	start := 0

	if len(q.columns) == 0 {
		q.columns = make([]string, len(rows[0]))
		for i, value := range rows[0] {
			q.columns[i] = value.(string)
		}

		start = 1
	}

	// Fill records maps
	for _, row := range rows[start:] {
		r := make(Record)
		for i, value := range row {
			r.set(q.columns[i], value)
		}
		records = append(records, r)
	}

	return records, nil
}

func newQuery(table string, ls *Livestatus) *Query {
	return &Query{
		table:   table,
		headers: make([]string, 0),
		ls:      ls,
	}
}
