package cmkapi

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Client holds the struct for connecting to a Check_MK instance
type Client struct {
	User       string
	Password   string
	Host       string
	httpClient *http.Client
}

// NewClient is the skeleton for instantiating a new Check_MK connection
func NewClient(user, password, host string) (*Client, error) {
	return &Client{user, password, host, nil}, nil
}

// NewAPIRequest sends and receives the Check_MK webAPI
func (c *Client) NewAPIRequest(method, APICall string, body io.Reader) (resp_body []byte, resp_error error) {
	baseurl := c.Host
	action := "?action=" + APICall
	credentails := "&_username=" + c.User + "&_secret=" + c.Password + "&effective_attributes=1"
	apiurl := baseurl + action + credentails
	request, requestErr := http.NewRequest("POST", apiurl, body)
	if requestErr != nil {
		log.Fatalf("ERROR: %s", requestErr)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, doErr := (&http.Client{}).Do(request)
	if doErr != nil {
		log.Fatalf("ERROR: %s", doErr)
	}
	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	return respBody, err

}
