package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccDatasourceRollbarUser_Basic(t *testing.T) {
	userID := testAccConfig.GetUserOrAbort(t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarUserWithDatasourceBasic(userID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.rollbar_user.foobar", "id", userID),
				),
			},
		},
	})
}

func testAccCheckRollbarUserWithDatasourceBasic(userID string) string {
	return fmt.Sprintf(`
data "rollbar_user" "foobar" {
  id = "%s"
}
`, userID)
}
