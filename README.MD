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
                //client.ReadHost("winxp_1")
                client.ReadALLHost1()
                //client.DeleteHost("winxp_1")

        }
}

```
