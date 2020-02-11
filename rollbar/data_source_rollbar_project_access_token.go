package rollbar

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceRollbarProjectAccessToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRollbarProjectAccessTokenRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"access_tokens": {
				Type:      schema.TypeMap,
				Sensitive: true,
				Computed:  true,
				Elem:      schema.TypeString,
			},
		},
	}
}

func dataSourceRollbarProjectAccessTokenRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(GenerateRandomResourceID())

	client := m.(*Config).API
	projectID := getProjectID(d)

	result, _, getErr := client.ProjectAccessTokens.List(projectID)
	if getErr != nil {
		return getErr
	}

	// Loop through and create a map of string:string, where the key is the token name
	// and the value is the access token.
	tokenMap := make(map[string]string)
	for _, pat := range result.Result {
		// Only store enabled tokens
		if pat.GetStatus() == "enabled" {
			tokenMap[pat.GetName()] = pat.GetAccessToken()
		}
	}

	return d.Set("access_tokens", tokenMap)
}
