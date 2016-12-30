package cmkapi

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)



func (session *Session) GetHosts() error {
	var body []byte
	var response *http.Response
	var request *http.Request
	fullurl := session.BaseURL + "?action=get_all_hosts&_username=" + session.Username + "&_secret=" + session.Password
	//fmt.Printf("%s", fullurl)

	request, err := http.NewRequest("GET", fullurl, nil)
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
