package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccRollbarProjectAccessToken_importBasic(t *testing.T) {
	projectName := fmt.Sprintf("project-tftest-%s", acctest.RandString(10))
	tokenName := fmt.Sprintf("token-tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarProjectAccessToken_basic(projectName, tokenName),
			},
			{
				ResourceName:      "rollbar_project_access_token.foobar",
				ImportStateIdFunc: testAccRollbarProjectAccessTokenImportStateIdFunc("rollbar_project_access_token.foobar"),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccRollbarProjectAccessTokenImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("%s:%s", rs.Primary.Attributes["project_id"],
			rs.Primary.Attributes["access_token"]), nil
	}
}
