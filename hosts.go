package cmkapi

import (
    "fmt"
    "strings"
    "io/ioutil"
    "log"
    "net/http"
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



func (c *Client) ReadALLHost() error {
	baseurl := "http://" + c.Host + "/" + c.Sitename + "/check_mk/webapi.py"
	action := "?action=get_all_hosts"
	//action := "?action=get_host"
	credentails := "&_username=" + c.User + "&_secret=" + c.Password
	fullurl := baseurl + action + credentails 

	request, err := http.NewRequest("GET", fullurl, nil)
	if err == nil {
                request.Header.Add("Content-Type", "application/json")
                response, err1 := (&http.Client{}).Do(request)
		if err1 == nil {
                	defer response.Body.Close()
                	body, err2 := ioutil.ReadAll(response.Body)
			if err2 == nil {
                		fmt.Printf("%s", body)
			} else {
			log.Fatalf("ERROR: %s", err2)
			}
		} else { 
		log.Fatalf("ERROR: %s", err1)
        }
		
        } else {
                log.Fatalf("ERROR: %s", err)
        }
	return nil
}

func (c *Client) ReadHost(host string) error {
	baseurl := "http://" + c.Host + "/" + c.Sitename + "/check_mk/webapi.py"
	action := "?action=get_host"
	credentails := "&_username=" + c.User + "&_secret=" + c.Password
	fullurl := baseurl + action + credentails
	s := "request={\"hostname\": \"" + host + "\"}"
	search := strings.NewReader(s)
	request, err := http.NewRequest("POST",fullurl,search)
	if err == nil {
                request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
                response, err1 := (&http.Client{}).Do(request)
                if err1 == nil {
                        defer response.Body.Close()
                        body, err2 := ioutil.ReadAll(response.Body)
                        if err2 == nil {
                                fmt.Printf("%s", body)
                        } else {
                        log.Fatalf("ERROR: %s", err2)
                        }
                } else {
                log.Fatalf("ERROR: %s", err1)
        }

        } else {
                log.Fatalf("ERROR: %s", err)
        }
        return nil

	
}
