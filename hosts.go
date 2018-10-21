package cmkapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CreateHost creates a new host
func (c *Client) CreateHost(hostname, folder, alias, tag_agent, tag_criticality, ipaddress string) error {
	var result StructPutResult
	host := &Host{Attributes{alias, tag_agent, tag_criticality, ipaddress}, hostname, folder}
	h, marshal_err := json.Marshal(host)
	if marshal_err != nil {
		fmt.Printf("Error Creating the Host Struct: %s\n", marshal_err)
		return marshal_err
	}
	s := "request=" + string(h)
	body := strings.NewReader(s)
	resp_body, resp_err := c.NewAPIRequest("POST", "add_host", body)
	if resp_err != nil {
		fmt.Printf("API Request for add_host failed. Error: %s\n", resp_err)
		return resp_err
	}
	resp_unmarshal_err := json.Unmarshal(resp_body, &result)
	if resp_unmarshal_err != nil {
		fmt.Printf("Error Decoding the API response. Error: %s\n", resp_unmarshal_err)
		return resp_unmarshal_err
	}
	resp_code := result.ResultCode
	if resp_code != 0 {
		fmt.Printf("API Response Failed. Error:%s\n", result.Result)
		return errors.New("API Response Failed")
	}
	//Add host is sucssfull , now Call activate_changes
	status := c.ActivateChanges()
	if status != nil {
		fmt.Printf("Activate Change got Failed, %s\n", status)
		return errors.New("Activate Change got Failed, status")
	}
	fmt.Printf("Host %s is added Successfully \n", host.Hostname)
	return nil
}

// ReadALLHost returns all hosts
func (c *Client) ReadALLHost() error {
	resp_body, resp_err := c.NewAPIRequest("GET", "get_all_hosts", nil)
	if resp_err != nil {
		return resp_err
	}
	fmt.Printf("%s", resp_body)
	return nil
}

// ReadHost returns a single host
func (c *Client) ReadHost(host string) (*Host, error) {
	var hostdetail StructGetHostResult
	s := "request={\"hostname\": \"" + host + "\"}"
	body := strings.NewReader(s)
	resp_body, resp_err := c.NewAPIRequest("POST", "get_host", body)
	if resp_err != nil {
		fmt.Printf("API Request for get_host failed. Error: %s\n", resp_err)
		return nil, resp_err
	}
	resp_unmarshal_err := json.Unmarshal(resp_body, &hostdetail)
	if resp_unmarshal_err != nil {
		fmt.Printf("Error Decoding the API response. Error: %s\n", resp_unmarshal_err)
		return nil, resp_unmarshal_err
	}
	hostname := hostdetail.Result.Hostname
	folder := hostdetail.Result.Path
	alias := hostdetail.Result.Attributes.Alias
	tag_agent := hostdetail.Result.Attributes.TagAgent
	tag_criticality := hostdetail.Result.Attributes.TagCriticality
	ipaddress := hostdetail.Result.Attributes.Ipaddress
	hoststruct := &Host{Attributes{alias, tag_agent, tag_criticality, ipaddress}, hostname, folder}
	fmt.Printf("Host %s is Available in Check_MK.\n", hostname)
	return hoststruct, nil
}

// DeleteHost deletes a host
func (c *Client) DeleteHost(host string) error {
	var result StructPutResult
	s := "request={\"hostname\": \"" + host + "\"}"
	body := strings.NewReader(s)
	resp_body, resp_err := c.NewAPIRequest("POST", "delete_host", body)
	if resp_err != nil {
		fmt.Printf("API Request for delete_host failed. Error: %s\n", resp_err)
		return resp_err
	}
	resp_unmarshal_err := json.Unmarshal(resp_body, &result)
	if resp_unmarshal_err != nil {
		fmt.Printf("Error Decoding the API response. Error: %s\n", resp_unmarshal_err)
		return resp_unmarshal_err
	}
	resp_code := result.ResultCode
	if resp_code != 0 {
		fmt.Printf("API Response Failed. Error:%s\n", result.Result)
		return errors.New("API Response Failed")
	}
	//Delete host is sucssfull , now Call activate_changes
	status := c.ActivateChanges()
	if status != nil {
		fmt.Printf("Activate Change got Failed, %s\n", status)
		return errors.New("Activate Change got Failed, status")
	}
	fmt.Printf("Host %s is deleted Successfully \n", host)
	return nil
}

// ActivateChanges activates the pending changes done to Check_MK after create/update/delete
func (c *Client) ActivateChanges() error {
	var result StructPutResult
	resp_body, resp_err := c.NewAPIRequest("POST", "activate_changes", nil)
	if resp_err != nil {
		fmt.Printf("API Request for activate_changes failed. Error: %s\n", resp_err)
		return resp_err
	}
	resp_unmarshal_err := json.Unmarshal(resp_body, &result)
	if resp_unmarshal_err != nil {
		fmt.Printf("Error Decoding the API response. Error: %s\n", resp_unmarshal_err)
		return resp_unmarshal_err
	}
	resp_code := result.ResultCode
	if resp_code != 0 {
		fmt.Printf("API Response Failed. Error:%s\n", result.Result)
		return errors.New("API Response Failed")
	}
	return nil
}
