package simpleresty

import "github.com/go-resty/resty/v2"

// New function creates a new SimpleResty client.
func New() *Client {
	return &Client{Client: resty.New()}
}
