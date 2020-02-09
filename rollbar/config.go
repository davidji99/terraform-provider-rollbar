package rollbar

import (
	"encoding/json"
	"github.com/davidji99/terraform-provider-rollbar/rollbar_api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Config struct {
	API                *rollbar_api.Client
	Headers            map[string]string
	accountAccessToken string
	projectAccessToken string
}

func NewConfig() *Config {
	config := &Config{}
	return config
}

func (c *Config) initializeAPI() error {
	authConfig := &rollbar_api.TokenAuthConfig{
		AccountAccessToken: &c.accountAccessToken,
		CustomHttpHeaders:  c.Headers,
	}

	api, clientInitErr := rollbar_api.NewClientTokenAuth(authConfig)
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
