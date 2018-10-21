package cmkapi

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	User       string
	Password   string
	Host       string
	Sitename   string
	httpClient *http.Client
}

//#-------------------------------------------------------------------------------------------------------------------------------------------
// func NewClient
func NewClient(user, password, host, sitename string) (*Client, error) {
	return &Client{user, password, host, sitename, nil}, nil
}

//#-------------------------------------------------------------------------------------------------------------------------------------------
func (c *Client) NewAPIRequest(method, APICall string, body io.Reader) (resp_body []byte, resp_error error) {
	baseurl := "http://" + c.Host + "/" + c.Sitename + "/check_mk/webapi.py"
	action := "?action=" + APICall
	credentails := "&_username=" + c.User + "&_secret=" + c.Password + "&effective_attributes=1"
	apiurl := baseurl + action + credentails
	request, request_err := http.NewRequest("POST", apiurl, body)
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
	return resp_body, err

}

//#-------------------------------------------------------------------------------------------------------------------------------------------
