package rollbar

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
import fmt "fmt"

func dataSourceRollbarProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRollbarProjectRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceRollbarProjectRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).API
	name := d.Get("name").(string)

	result, _, err := client.Projects.List()

	if err != nil {
		return err
	}

	if result.HasResults() {
		for _, project := range result.Results {
			if project.GetName() == name {
				d.SetId(Int64ToString(project.GetID()))

				var setErr error
				setErr = d.Set("status", project.GetStatus())
				setErr = d.Set("account_id", project.GetAccountID())
				setErr = d.Set("name", project.GetName())
				return setErr
			}
		}
	}

	return fmt.Errorf("no matches found for project name: %s", name)
}
