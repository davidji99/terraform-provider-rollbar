package simpleresty

import (
	"github.com/go-resty/resty/v2"
	"os"
)

var (
	proxyVars = []string{"HTTPS_PROXY", "https_proxy", "HTTP_PROXY", "http_proxy"}
)

// New function creates a new SimpleResty client.
func New() *Client {
	c := &Client{Client: resty.New()}

	determineSetProxy(c)

	return c
}

// determineSetProxy checks if any proxy variables are defined in the environment.
// If so, set the first occurrence and exit the loop.
func determineSetProxy(c *Client) {
	for _, v := range proxyVars {
		proxyUrl := os.Getenv(v)
		if proxyUrl != "" {
			c.SetProxy(proxyUrl)
			break
		}
	}
}
