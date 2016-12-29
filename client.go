package cmkapi

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type Session struct {
        Username           string
        Password           string
        BaseURL            string
        AllowUnverifiedSSL bool
        httpClient         *http.Client
}

// func NewSession ...
func NewSession(username, password, url string,allowUnverifiedSSL bool) (*Session, error) {
        return &Session{username, password, url,allowUnverifiedSSL, nil}, nil
}

func (session *Session) Connect() error {
	var body []byte
	var response *http.Response
	var request *http.Request
	request, err := http.NewRequest("GET", session.BaseURL, nil)
    	if err == nil {
        	request.Header.Add("Content-Type", "application/json")
        	response, err = (&http.Client{}).Do(request)
        	defer response.Body.Close()
        	body, err = ioutil.ReadAll(response.Body)
        	fmt.Printf("%s", body)
    	} else {
        	log.Fatalf("ERROR: %s", err)
    	}
	return nil
}
