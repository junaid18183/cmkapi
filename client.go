package cmkapi

import ( "net/http" )

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

