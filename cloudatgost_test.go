package cloudatgost

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func testTools(code int, body string) (*httptest.Server, *Client)  {
  server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(code)
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, body)
  }))

  // transport := &http.Transport{
  //   Proxy: func(req *http.Request) (*url.URL, error) {
  //     return url.Parse(server.URL)
  //   },
  // }
  // httpClient := &http.Client{Transport: transport}
  httpClient := http.DefaultClient

  client := NewClient("johndoe@example.com", "myApiKey", httpClient)
  client.BaseURL = server.URL

  return server, client
}


func TestNewClient (t *testing.T) {
	server, client := testTools(200, "myResponseBody")
	defer server.Close()

	if(client.Login != "johndoe@example.com" && client.Token != "myApiKey") {
		t.Fatalf("Failed to properly instantiate a new client.")
	} else {
		t.Logf("Successfully instantiated a new client.")
	}
}

func TestListServers (t *testing.T) {
	server, client := testTools(200, "{\"status\":\"ok\",\"time\":1425064819,\"id\":\"90000\",\"data\":[{\"id\":\"2402939\",\"CustID\":\"90001\",\"packageid\":\"21579\",\"servername\":\"localhost\",\"label\":\"api-test\",\"vmname\":\"c90001-iOD-848113\",\"ip\":\"1.2.3.4\",\"netmask\":\"255.255.255.0\",\"gateway\":\"1.2.3.1\",\"hostname\":\"c9000-iOD-123.cloudatcost.com\",\"rootpass\":\"password\",\"vncport\":\"45148\",\"vncpass\":\"vnc_pass\",\"servertype\":\"Custom\",\"template\":\"Microsoft Windows 7 (64-bit)\",\"cpu\":\"4\",\"cpuusage\":\"261\",\"ram\":\"4096\",\"ramusage\":\"3556.202\",\"storage\":\"69\",\"hdusage\":\"0.000434688\",\"sdate\":\"01\\/26\\/2014\",\"status\":\"Powered On\",\"panel_note\":\"\",\"mode\":\"Normal\",\"uid\":\"619712052\",\"sid\":\"240294\"}]}")
	defer server.Close()

	servers := client.ListServers()

	if (servers.Status != "ok" && servers.Data[0].ID != "2402939") {
		t.Fatalf("placeholder did not work as expected.")
	} else {
		t.Logf("JSON output successfully deserialized to ServerList.")
	}
}

func TestListTemplates (t *testing.T) {
	server, client := testTools(200, "{\"status\":\"ok\",\"time\":1425326406,\"data\":[{\"id\":\"26\",\"detail\":\"CentOS-7-64bit\"},{\"id\":\"27\",\"detail\":\"Ubuntu-14.04.1-LTS-64bit\"},{\"id\":\"15\",\"detail\":\"CentOS 6.5 64bit (LAMP)\"},{\"id\":\"21\",\"detail\":\"Ubuntu 12.10 64bit\"},{\"id\":\"23\",\"detail\":\"Ubuntu 12.04.3 LTS 64bit\"},{\"id\":\"24\",\"detail\":\"Windows 2008 R2 64bit (BigDogs Only)\"},{\"id\":\"25\",\"detail\":\"Windows 2012 R2 64bit (BigDogs Only)\"},{\"id\":\"14\",\"detail\":\"CentOS 6.5 64bit (cPanel-WHM)\"},{\"id\":\"13\",\"detail\":\"CentOS 6.5 64bit\"},{\"id\":\"10\",\"detail\":\"CentOS 6.5 32bit\"},{\"id\":\"3\",\"detail\":\"Debian 7.1 64bit\"},{\"id\":\"9\",\"detail\":\"Windows7 64bit (BigDogs Only)\"},{\"id\":\"2\",\"detail\":\"Ubuntu-13.10-64bit\"},{\"id\":\"1\",\"detail\":\"CentOS 6.4 64bit\"},{\"id\":\"28\",\"detail\":\"Minecraft-CentOS-7-64bit\"}]}")
	defer server.Close()

	templates := client.ListTemplates()

	if (templates.Status != "ok" && templates.Data[0].Detail != "CentOS-7-64bit") {
		t.Fatalf("placeholder did not work as expected.")
	} else {
		t.Logf("JSON output successfully deserialized to TemplateList.")
	}
}

func TestListTasks (t *testing.T) {
	server, client := testTools(200, "{\"status\":\"ok\",\"time\":1425504688,\"api\":\"v1\",\"cid\":\"734103810\",\"action\":\"listtasks\",\"data\":[{\"cid\":\"734103810\",\"idf\":\"8548136390745\",\"serverid\":\"0\",\"action\":\"reset\",\"status\":\"completed\",\"starttime\":\"1425504093\",\"finishtime\":\"1425504094\"},{\"cid\":\"734103810\",\"idf\":\"2268428551033\",\"serverid\":\"254513205\",\"action\":\"reset\",\"status\":\"pending\",\"starttime\":\"1425504295\",\"finishtime\":\"1425504312\"}]}")
	defer server.Close()

	tasks := client.ListTasks()

	if (tasks.Status != "ok" && tasks.Data[0].Action != "reset") {
		t.Fatalf("placeholder did not work as expected.")
	} else {
		t.Logf("JSON output successfully deserialized to TaskList.")
	}
}

func TestConsole (t *testing.T) {
	server, client := testTools(200, "{\"status\":\"ok\",\"time\":1425572027,\"api\":\"v1\",\"serverid\":\"1234567890\",\"console\":\"http:\\/\\/panel.cloudatcost.com:12345\\/console.html?servername=123456&hostname=1.1.1.1&sshkey=123456&sha1hash=aBcDeFgG\"}")
	defer server.Close()

	console := client.Console("1234")

	if (console.Status != "ok" && console.Console != "http:\\/\\/panel.cloudatcost.com:12345\\/console.html?servername=123456&hostname=1.1.1.1&sshkey=123456&sha1hash=aBcDeFgG") {
		t.Fatalf("placeholder did not work as expected.")
	} else {
		t.Logf("JSON output successfully deserialized to Console.")
	}
}

func TestAction (t *testing.T) {
	server, client := testTools(200, "{\"status\":\"ok\",\"time\":1425504815,\"api\":\"v1\",\"serverid\":\"254513205\",\"action\":\"poweron\",\"taskid\":700420024805,\"result\":\"successful\"}")
	defer server.Close()

	op := client.Action("1234", "poweron")

	if (op.Status != "ok" && op.Action != "poweron") {
		t.Fatalf("placeholder did not work as expected.")
	} else {
		t.Logf("JSON output successfully deserialized to Action.")
	}
}

func TestPowerOn (t *testing.T) {
	server, client := testTools(200, "{\"status\":\"ok\",\"time\":1425504815,\"api\":\"v1\",\"serverid\":\"254513205\",\"action\":\"poweron\",\"taskid\":700420024805,\"result\":\"successful\"}")
	defer server.Close()

	op := client.PowerOn("1234")

	if (op.Status != "ok" && op.Action != "poweron") {
		t.Fatalf("placeholder did not work as expected.")
	} else {
		t.Logf("JSON output successfully deserialized to Action.")
	}
}

func TestPowerOff (t *testing.T) {
	server, client := testTools(200, "{\"status\":\"ok\",\"time\":1425504815,\"api\":\"v1\",\"serverid\":\"254513205\",\"action\":\"poweroff\",\"taskid\":700420024805,\"result\":\"successful\"}")
	defer server.Close()

	op := client.PowerOff("1234")

	if (op.Status != "ok" && op.Action != "poweroff") {
		t.Fatalf("placeholder did not work as expected.")
	} else {
		t.Logf("JSON output successfully deserialized to Action.")
	}
}

func TestReboot (t *testing.T) {
	server, client := testTools(200, "{\"status\":\"ok\",\"time\":1425504815,\"api\":\"v1\",\"serverid\":\"254513205\",\"action\":\"reset\",\"taskid\":700420024805,\"result\":\"successful\"}")
	defer server.Close()

	op := client.Reboot("1234")

	if (op.Status != "ok" && op.Action != "reset") {
		t.Fatalf("placeholder did not work as expected.")
	} else {
		t.Logf("JSON output successfully deserialized to Action.")
	}
}
