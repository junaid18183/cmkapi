This is sample API for Check_MK

example 
```
package main
import (
        "github.com/junaid18183/cmkapi"
//      "fmt"
)

func main() {

        // account credentials and url to connect to grafana
        user := "autouser"
        pass := "UPFKWAJJDPJWTOQMOWHY"
        host := "192.168.99.100:32768"
        sitename := "mva"
        client,error := cmkapi.NewClient(user,pass,host,sitename)

        if error == nil {
                //client.CreateHost("HostName", "Folder","Alias of host", "Agent-Type{cmk-agent,snmp-only,snmp-v1,snmp-tcp,ping}","Criticality-Tag{prod,critical,test,offline}","IP Address")
                client.CreateHost("winxp_1", "os/windows","Alias of winxp_1", "cmk-agent","prod","127.0.0.1")
                client.ReadHost("winxp_1")
                client.DeleteHost("winxp_1")

        }
}

```
