package cmkapi

import (
    "fmt"
    "strings"
    "io/ioutil"
    "log"
    "net/http"
    "encoding/json"
)

type Host struct {
	Attributes
        Hostname string
        Folder   string
}

type Attributes struct {
	Alias string
	Agent string
	IPADDRESS string
	Criticality string
}


func (h *Host) Id() string {
        return "id-" + h.Hostname + "!"
}

func (c *Client) CreateHost(h *Host) error {
        return nil
}

//#-------------------------------------------------------------------------------------------------------------------------------------------
func (c *Client) ReadALLHost() error {
	resp_body, resp_err := c.NewAPIRequest("GET","get_all_hosts",nil)
	if resp_err != nil {
        	return resp_err
        }
        fmt.Printf("%s",resp_body)
        return nil

}
//#-------------------------------------------------------------------------------------------------------------------------------------------
func (c *Client) ReadHost(host string) error {
        var hostdetail StructGetHostResult
        s := "request={\"hostname\": \"" + host + "\"}"
        body := strings.NewReader(s)
	resp_body, resp_err := c.NewAPIRequest("POST","get_host",body)
	if resp_err != nil {
		return resp_err
	}
	err := json.Unmarshal(resp_body, &hostdetail)
        if err == nil {
        fmt.Printf("%s", hostdetail.Result.Hostname)
        }
        return nil
}
//#-------------------------------------------------------------------------------------------------------------------------------------------
func (c *Client) DeleteHost(host string) error {
        s := "request={\"hostname\": \"" + host + "\"}"
        body := strings.NewReader(s)
	resp_body, resp_err := c.NewAPIRequest("POST","delete_host",body)
        if resp_err != nil {
		return resp_err
	}
	fmt.Printf("%s",resp_body)
        return nil
}
//#-------------------------------------------------------------------------------------------------------------------------------------------
