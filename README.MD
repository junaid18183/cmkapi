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
        url := "http://192.168.99.100:32768/mva/check_mk/webapi.py"
        session,error := cmkapi.NewSession(user,pass,url,true)
        if error == nil {
                session.GetHosts()
                //fmt.Println("Done")
        }
}

```
