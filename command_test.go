package livestatus

import (
	"testing"
	"time"
)

//ACKNOWLEDGE_HOST_PROBLEM;<host_name>;<sticky>;<notify>;<persistent>;<author>;<comment>
//ACKNOWLEDGE_SVC_PROBLEM;<host_name>;<service_description>;<sticky>;<notify>;<persistent>;<author>;<comment>
//SCHEDULE_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
//SCHEDULE_SVC_DOWNTIME;<host_name>;<service_desription><start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>

func Test_Cmd_AcknowledgeHostProblem(t *testing.T) {
	expected := "COMMAND 12345678 ACKNOWLEDGE_HOST_PROBLEM;host;2;1;1;me@example.com;some stuff\n\n"

	c := newCommand(&Livestatus{}, "ACKNOWLEDGE_HOST_PROBLEM")
	c.Hostname("host")
	c.Sticky(true)
	c.Notify(true)
	c.Persistent(true)
	c.Author("me@example.com")
	c.Comment("some stuff")

	result, err := c.buildCmd(time.Unix(12345678, 0))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}

	if result != expected {
		t.Logf("\nExpected %q\nbut got  %q\n", expected, result)
		t.Fail()
	}
}
