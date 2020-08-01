package rollbar

import (
	"fmt"
	"github.com/davidji99/rollrest-go/rollrest"
	"github.com/davidji99/terraform-provider-rollbar/version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Config struct {
	API                                       *rollrest.Client
	Headers                                   map[string]string
	accountAccessToken                        string
	projectAccessToken                        string
	PostCreatePDIntegrationDeleteDefaultRules bool
}

func NewConfig() *Config {
	config := &Config{}
	return config
}

func (c *Config) initializeAPI() error {
	userAgent := fmt.Sprintf("terraform-provider-rollbar/v%s", version.ProviderVersion)

	api, clientInitErr := rollrest.New(rollrest.AuthAAT(c.accountAccessToken), rollrest.AuthPAT(c.projectAccessToken),
		rollrest.CustomHTTPHeaders(c.Headers), rollrest.UserAgent(userAgent))
	if clientInitErr != nil {
		return clientInitErr
	}
	c.API = api

	return nil
}

func (c *Config) applySchema(d *schema.ResourceData) (err error) {
	if v, ok := d.GetOk("headers"); ok {
		headersRaw := v.(map[string]interface{})
		h := make(map[string]string)

		for k, v := range headersRaw {
			h[k] = fmt.Sprintf("%v", v)
		}

		c.Headers = h
	}

	c.PostCreatePDIntegrationDeleteDefaultRules = d.Get("post_create_pd_integration_delete_default_rules").(bool)

	return nil
}
