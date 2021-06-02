package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccRollbarTeamUserAssociationTest_importBasic(t *testing.T) {
	teamID := testAccConfig.GetTeamIDorAbort(t)
	email := testAccConfig.GetTeamEmailAddress(t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarTeamUserAssociation_basic(teamID, email),
			},
			{
				ResourceName:      "rollbar_team_user_association.foobar",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccRollbarTeamUserAssociationImportStateIdFunc("rollbar_team_user_association.foobar"),
			},
		},
	})
}

func testAccRollbarTeamUserAssociationImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("%s:%s", rs.Primary.Attributes["team_id"],
			rs.Primary.Attributes["email"]), nil
	}
}
