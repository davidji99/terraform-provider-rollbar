package rollbar

import (
	"encoding/json"
	"github.com/davidji99/terraform-provider-rollbar/rollapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Config struct {
	API                *rollapi.Client
	Headers            map[string]string
	accountAccessToken string
	projectAccessToken string
}

func NewConfig() *Config {
	config := &Config{}
	return config
}

func (c *Config) initializeAPI() error {
	authConfig := &rollapi.TokenAuthConfig{
		AccountAccessToken: &c.accountAccessToken,
		CustomHttpHeaders:  c.Headers,
	}

	api, clientInitErr := rollapi.NewClientTokenAuth(authConfig)
	if clientInitErr != nil {
		return clientInitErr
	}
	c.API = api

	return nil
}

func (c *Config) applySchema(d *schema.ResourceData) (err error) {
	headers := make(map[string]string)
	if h := d.Get("api_headers").(string); h != "" {
		if err = json.Unmarshal([]byte(h), &headers); err != nil {
			return
		}
		c.Headers = headers
	}

	return nil
}
