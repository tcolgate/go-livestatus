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

type argDef struct {
	t      string
	fmtStr string
}

var argTypes = map[string]argDef{
	"host_name":               {"string", `Hostname(%s)`},
	"service_description":     {"string", `ServiceDescription(%s)`},
	"sticky":                  {"bool", `Sticky(%s)`},
	"notify":                  {"bool", `Notify(%s)`},
	"fixed":                   {"bool", `Fixed(%s)`},
	"persistent":              {"bool", "Persistent(%s)"},
	"author":                  {"string", `Author(%s)`},
	"contact_name":            {"string", `ContactName(%s)`},
	"comment":                 {"string", `Comment(%s)`},
	"start_time":              {"time.Time", `Start(%s)`},
	"end_time":                {"time.Time", `End(%s)`},
	"notification_timeperiod": {"string", `NotificationTimePeriod(%s)`},
	"duration":                {"time.Duration", `Duration(%s)`},
	"trigger_id":              {"int", `TriggerID(%s)`},
	"downtime_id":             {"int", `DowntimeID(%s)`},
}

func main() {
	client := http.Client{}
	buf := &bytes.Buffer{}
	resp, err := client.Get(nagCmdIndex)
	if err != nil {
		log.Fatalf("error fetching index, %v", err.Error())
	}

	io.Copy(buf, resp.Body)

	cs := nagCmdURLRgxp.FindAllStringSubmatch(buf.String(), -1)
	for _, cUrl := range cs {
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

		//if i > 2 {
		//		break
		//	}
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
		args := []string{}
		for _, a := range c.args {
			args = append(args, fmt.Sprintf("%s %s", a, argTypes[a].t))
		}
		fmt.Printf("func (c *Command) %v(%s){\n", goname, strings.Join(args, ", "))
		fmt.Printf("\tc.Raw(\"%v\")\n", c.name)
		for _, a := range c.args {
			fmt.Printf("\tc.%s\n", fmt.Sprintf(argTypes[a].fmtStr, a))
		}
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
