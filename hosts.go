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



func (c *Client) ReadALLHost() error {
	fullurl:= c.APIURL("get_all_hosts") 

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
	var hostdetail GetHostResult
        fullurl:= c.APIURL("get_host")
	s := "request={\"hostname\": \"" + host + "\"}"
        search := strings.NewReader(s)
        request, err := http.NewRequest("POST",fullurl,search)
        if err == nil {
                request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
                response, doerr := (&http.Client{}).Do(request)
                if doerr == nil {
                        defer response.Body.Close()
                        body, err2 := ioutil.ReadAll(response.Body)
                        if err2 == nil {
                                //fmt.Printf("%s", body)
				err := json.Unmarshal(body, &hostdetail)
				if err == nil {
                                fmt.Printf("%s", hostdetail.Result.Hostname)
				}
				
                        } else {
                        log.Fatalf("ERROR: %s", err2)
                        }
                } else {
                log.Fatalf("ERROR: %s", doerr)
        }

        } else {
                log.Fatalf("ERROR: %s", err)
        }
        return nil


}

func (c *Client) DeleteHost(host string) error {
        fullurl:= c.APIURL("delete_host")
        s := "request={\"hostname\": \"" + host + "\"}"
        body := strings.NewReader(s)
        request, request_err := http.NewRequest("POST",fullurl,body)
	if request_err != nil {
		log.Fatalf("ERROR: %s", request_err) 
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, do_err := (&http.Client{}).Do(request)
	if do_err != nil {
		log.Fatalf("ERROR: %s", do_err)
	}
	defer response.Body.Close()
	resp_body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("%s -  %s", resp_body, err)
        return nil

}

