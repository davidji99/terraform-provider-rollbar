package rollbar

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"project_access_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROLLBAR_PROJECT_ACCESS_TOKEN", nil),
			},

			"account_access_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROLLBAR_ACCOUNT_ACCESS_TOKEN", nil),
			},

			"api_headers": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROLLBAR_API_HEADERS", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"rollbar_project":              resourceRollbarProject(),
			"rollbar_project_access_token": resourceRollbarProjectAccessToken(),
			"rollbar_team":                 resourceRollbarTeam(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"rollbar_team": dataSourceRollbarTeam(),
		},

		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure configures the rollbar api client
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Println("[INFO] Initializing Rollbar Provider")

	config := NewConfig()

	if accountAccessToken, ok := d.GetOk("account_access_token"); ok {
		config.accountAccessToken = accountAccessToken.(string)
	}

	if projectAccessToken, ok := d.GetOk("project_access_token"); ok {
		config.projectAccessToken = projectAccessToken.(string)
	}

	if applySchemaErr := config.applySchema(d); applySchemaErr != nil {
		return nil, applySchemaErr
	}

	if err := config.initializeAPI(); err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Rollbar provider initialized")

	return config, nil
}
