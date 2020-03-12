package rollbar

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceRollbarProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRollbarProjectRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
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
				d.Set("status", project.GetStatus())
				d.Set("account_id", project.GetAccountID())
				return nil
			}
		}
	}

	d.SetId("")
	return nil
}
