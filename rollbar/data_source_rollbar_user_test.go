package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDatasourceRollbarUser_Basic(t *testing.T) {
	email := testAccConfig.GetUserEmailOrAbort(t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarUserWithDatasourceBasic(email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.rollbar_user.foobar", "email", email),
					resource.TestCheckResourceAttrSet("data.rollbar_user.foobar", "username"),
				),
			},
		},
	})
}

func testAccCheckRollbarUserWithDatasourceBasic(email string) string {
	return fmt.Sprintf(`
data "rollbar_user" "foobar" {
  email = "%s"
}
`, email)
}
