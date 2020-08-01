package rollbar

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

			"headers": {
				Type:     schema.TypeMap,
				Elem:     schema.TypeString,
				Optional: true,
			},

			"post_create_pd_integration_delete_default_rules": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"rollbar_project":               dataSourceRollbarProject(),
			"rollbar_project_access_tokens": dataSourceRollbarProjectAccessTokens(),
			"rollbar_team":                  dataSourceRollbarTeam(),
			"rollbar_user":                  dataSourceRollbarUser(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"rollbar_pagerduty_integration":       resourceRollbarPagerDutyIntegration(),
			"rollbar_pagerduty_notification_rule": resourceRollbarPagerDutyNotificationRule(),
			"rollbar_project":                     resourceRollbarProject(),
			"rollbar_project_access_token":        resourceRollbarProjectAccessToken(),
			"rollbar_team":                        resourceRollbarTeam(),
		},

		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure configures the rollbar api client
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Println("[INFO] Initializing Rollbar Provider")

	config := NewConfig()

	if accountAccessToken, ok := d.GetOk("account_access_token"); ok {
		log.Printf("[DEBUG] account_access_token to be used: %v", accountAccessToken)
		config.accountAccessToken = accountAccessToken.(string)
	}

	if projectAccessToken, ok := d.GetOk("project_access_token"); ok {
		log.Printf("[DEBUG] project_access_token to be used: %v", projectAccessToken)
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
