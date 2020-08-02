package rollbar

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRollbarProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRollbarProjectRead,
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

func dataSourceRollbarProjectRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Config).API
	name := d.Get("name").(string)

	result, _, err := client.Projects.List()

	if err != nil {
		diag.FromErr(err)
	}

	if result.HasResult() {
		for _, project := range result.Result {
			if project.GetName() == name {
				d.SetId(Int64ToString(project.GetID()))

				var diags diag.Diagnostics
				SetAttribute(d, diags, "status", project.GetStatus())
				SetAttribute(d, diags, "account_id", project.GetAccountID())
				SetAttribute(d, diags, "name", project.GetName())

				return diags
			}
		}
	}

	return diag.Errorf("no matches found for project name: %s", name)
}
