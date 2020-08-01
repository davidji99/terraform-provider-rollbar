package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDatasourceRollbarProjectAccessToken_Basic(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarProjectAccessTokenWithDatasourceBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.rollbar_project_access_token.foobar", "access_tokens.post_client_item"),
					resource.TestCheckResourceAttrSet(
						"data.rollbar_project_access_token.foobar", "access_tokens.post_server_item"),
					resource.TestCheckResourceAttrSet(
						"data.rollbar_project_access_token.foobar", "access_tokens.read"),
					resource.TestCheckResourceAttrSet(
						"data.rollbar_project_access_token.foobar", "access_tokens.write"),
				),
			},
		},
	})
}

func testAccCheckRollbarProjectAccessTokenWithDatasourceBasic(projectName string) string {
	return fmt.Sprintf(`
resource "rollbar_project" "foobar" {
	name = "%s"
}

data "rollbar_project_access_token" "foobar" {
  project_id = rollbar_project.foobar.id
}
`, projectName)
}
