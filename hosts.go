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
	h, marshalErr := json.Marshal(host)
	if marshalErr != nil {
		fmt.Printf("Error Creating the Host Struct: %s\n", marshalErr)
		return marshalErr
	}
	s := "request=" + string(h)
	body := strings.NewReader(s)
	respBody, respErr := c.NewAPIRequest("POST", "add_host", body)
	if respErr != nil {
		fmt.Printf("API Request for add_host failed. Error: %s\n", respErr)
		return respErr
	}
	respUnmarshalErr := json.Unmarshal(respBody, &result)
	if respUnmarshalErr != nil {
		fmt.Printf("Error Decoding the API response. Error: %s\n", respUnmarshalErr)
		return respUnmarshalErr
	}
	respCode := result.ResultCode
	if respCode != 0 {
		return errors.New("API Response Failed: " + result.Result)
	}
	//	//Add host is sucssfull , now Call activate_changes
	//	status := c.ActivateChanges()
	//	if status != nil {
	//		return errors.New("Activating Change(s) Failed: " + fmt.Sprintf("%s", status))
	//	}
	//	fmt.Printf("Host %s is added Successfully \n", host.Hostname)
	return nil
}

// ReadALLHost returns all hosts
func (c *Client) ReadALLHost() error {
	respBody, respErr := c.NewAPIRequest("GET", "get_all_hosts", nil)
	if respErr != nil {
		return respErr
	}
	fmt.Printf("%s", respBody)
	return nil
}

// ReadHost returns a single host
func (c *Client) ReadHost(host string) (*Host, error) {
	var hostdetail StructGetHostResult
	s := "request={\"hostname\": \"" + host + "\"}"
	body := strings.NewReader(s)
	respBody, respErr := c.NewAPIRequest("POST", "get_host", body)
	if respErr != nil {
		fmt.Printf("API Request for get_host failed. Error: %s\n", respErr)
		return nil, respErr
	}
	respUnmarshalErr := json.Unmarshal(respBody, &hostdetail)
	if respUnmarshalErr != nil {
		fmt.Printf("Error Decoding the API response. Error: %s\n", respUnmarshalErr)
		return nil, respUnmarshalErr
	}
	hostname := hostdetail.Result.Hostname
	folder := hostdetail.Result.Path
	alias := hostdetail.Result.Attributes.Alias
	tagAgent := hostdetail.Result.Attributes.TagAgent
	tagCriticality := hostdetail.Result.Attributes.TagCriticality
	ipaddress := hostdetail.Result.Attributes.Ipaddress
	hoststruct := &Host{Attributes{alias, tagAgent, tagCriticality, ipaddress}, hostname, folder}
	return hoststruct, nil
}

// DeleteHost deletes a host
func (c *Client) DeleteHost(host string) error {
	var result StructPutResult
	s := "request={\"hostname\": \"" + host + "\"}"
	body := strings.NewReader(s)
	respBody, respErr := c.NewAPIRequest("POST", "delete_host", body)
	if respErr != nil {
		fmt.Printf("API Request for delete_host failed. Error: %s\n", respErr)
		return respErr
	}
	respUnmarshalErr := json.Unmarshal(respBody, &result)
	if respUnmarshalErr != nil {
		fmt.Printf("Error Decoding the API response. Error: %s\n", respUnmarshalErr)
		return respUnmarshalErr
	}
	respCode := result.ResultCode
	if respCode != 0 {
		fmt.Printf("API Response Failed. Error:%s\n", result.Result)
		return errors.New("API Response Failed")
	}
	//Delete host is sucssfull , now Call activate_changes
	status := c.ActivateChanges()
	if status != nil {
		fmt.Printf("Activate Change got Failed, %s\n", status)

		return errors.New("Activating Change(s) Failed: " + fmt.Sprintf("%s", status))
	}
	fmt.Printf("Host %s is deleted Successfully \n", host)
	return nil
}

// ActivateChanges activates the pending changes done to Check_MK after create/update/delete
func (c *Client) ActivateChanges() error {
	var result StructPutResult
	respBody, respErr := c.NewAPIRequest("POST", "activate_changes", nil)
	if respErr != nil {
		fmt.Printf("API Request for activate_changes failed. Error: %s\n", respErr)
		return respErr
	}
	respUnmarshalErr := json.Unmarshal(respBody, &result)
	if respUnmarshalErr != nil {
		//fmt.Printf("Error Decoding the API response. Error: %s\n", respUnmarshalErr)
		return respUnmarshalErr
	}
	respCode := result.ResultCode
	if respCode != 0 {
		return errors.New("API Response Failed: " + result.Result)
	}
	return nil
}
