package cmkapi

import (
"net/http" 
)

type Client struct {
        User     string
        Password   string
        Host    string
        Sitename string
        httpClient         *http.Client
}


// func NewClient
func NewClient(user, password, host,sitename string) (*Client, error) {
        return &Client{user, password,host,sitename, nil}, nil
}


func (c *Client) APIURL(APICall string) (apiurl string) {
        baseurl := "http://" + c.Host + "/" + c.Sitename + "/check_mk/webapi.py"
        action := "?action=" + APICall
        credentails := "&_username=" + c.User + "&_secret=" + c.Password + "&effective_attributes=1"
        apiurl = baseurl + action + credentails
	return apiurl
}
