package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var nagCmdIndex = "https://www.nagios.org/developerinfo/externalcommands/"
var nagCmdURLRgxp = regexp.MustCompile(`<a href='(commandinfo.php\?command_id=[0-9]+)'>([A-Z_]+)</a></td></tr>`)

type nagCmd struct {
	def  string
	desc string

	name string
	args []string
}

var nagCmds = []nagCmd{}

func main() {
	client := http.Client{}
	buf := &bytes.Buffer{}
	resp, err := client.Get(nagCmdIndex)
	if err != nil {
		log.Fatalf("error fetching index, %v", err.Error())
	}

	io.Copy(buf, resp.Body)

	cs := nagCmdURLRgxp.FindAllStringSubmatch(buf.String(), -1)
	for i, cUrl := range cs {
		cmd := cUrl[2]
		ref, err := url.Parse(cUrl[1])
		if err != nil {
			log.Println("URL reference not valid, ", err.Error())
			continue
		}

		url := resp.Request.URL.ResolveReference(ref)
		cresp, err := client.Get(url.String())
		if err != nil {
			log.Fatalf("error fetching command %v, %v", cmd, err.Error())
		}

		doc, err := html.Parse(cresp.Body)
		if err != nil {
			log.Println("error parsing html, ", err.Error())
		}
		def, desc := findDefAndDesc(doc)
		nagCmds = append(nagCmds, nagCmd{def, desc, "", nil})

		if i > 2 {
			break
		}
	}

	genCode(os.Stdout, "livestatus", nagCmds)
}

func findDefAndDesc(n *html.Node) (string, string) {
	def := ""
	desc := ""
	foundDefHeader := false
	foundDescHeader := false

	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode && strings.Compare(n.Data, "Command Format:") == 0 {
			foundDefHeader = true
			return
		}
		if n.Type == html.TextNode && foundDefHeader && def == "" && strings.TrimSpace(n.Data) != "" {
			def = strings.TrimSpace(n.Data)
			return
		}

		if n.Type == html.TextNode && n.Data == "Description:" {
			foundDescHeader = true
			return
		}
		if n.Type == html.TextNode && foundDescHeader && desc == "" && strings.TrimSpace(n.Data) != "" {
			desc = strings.TrimSpace(n.Data)
			return
		}

		if def != "" && desc != "" {
			return
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return def, desc
}

func genCode(w io.Writer, pkg string, cmds []nagCmd) {
	fmt.Fprintf(w, "package %v\n\n", pkg)

	for _, c := range cmds {
		err := c.parseCommandDef()
		if err != nil {
			continue
		}
		nps := strings.Split(c.name, "_")
		goname := ""
		for _, np := range nps {
			goname = goname + strings.ToUpper(np[0:1]) + strings.ToLower(np[1:])
		}
		fmt.Printf("// %v is generated from the nagios external command definition:\n", goname)
		fmt.Printf("// Desfinition: %v\n", c.def)
		fmt.Println("// Description:")
		fmt.Printf("//  %v\n", c.desc)
		fmt.Printf("func %v(%#v){\n", goname, c.args)
		fmt.Printf("}\n\n")
	}
}

var cmdTmpl = regexp.MustCompile("^[A-Z_]+$")
var argTmpl = regexp.MustCompile("^<([a-z_]+)>$")

func (n *nagCmd) parseCommandDef() error {
	parts := strings.Split(n.def, ";")
	if !cmdTmpl.MatchString(parts[0]) {
		return fmt.Errorf("malformed command definition")
	}
	cmd := parts[0]

	args := []string{}
	for _, a := range parts[1:] {
		ms := argTmpl.FindStringSubmatch(a)
		if ms == nil {
			return fmt.Errorf("malformed command definition")
		}
		args = append(args, ms[1])
	}

	n.name = cmd
	n.args = args

	return nil
}

/*
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
*/
