package rollbar

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceRollbarUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRollbarUserRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRollbarUserRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("id").(string))

	client := m.(*Config).API

	user, _, getErr := client.Users.Get(StringToInt(d.Id()))
	if getErr != nil {
		return getErr
	}

	var setErr error
	setErr = d.Set("email", user.GetResult().GetEmail())
	setErr = d.Set("username", user.GetResult().GetUsername())

	return setErr
}
