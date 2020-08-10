package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceRollbarUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRollbarUserRead,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},

			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRollbarUserRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Config).API

	userEmail := d.Get("email").(string)

	users, _, getErr := client.Users.List()
	if getErr != nil {
		return getErr
	}

	for _, user := range users.GetResult().Users {
		if user.GetEmail() == userEmail {
			d.SetId(strconv.FormatInt(user.GetID(), 10))

			var setErr error
			setErr = d.Set("email", user.GetEmail())
			setErr = d.Set("username", user.GetUsername())

			return setErr
		}
	}

	return fmt.Errorf("could not find user %s in this account", userEmail)
}
