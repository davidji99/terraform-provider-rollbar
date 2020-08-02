package rollbar

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceRollbarTeam() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRollbarTeamRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"access_level": {
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

func dataSourceRollbarTeamRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("id").(string))

	return resourceRollbarTeamRead(d, m)
}
