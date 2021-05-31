package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccRollbarTeamProjectAssociation_importBasic(t *testing.T) {
	teamName := fmt.Sprintf("team-%s", acctest.RandString(10))
	projectName := fmt.Sprintf("project-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarTeamProjectAssociation_basic(teamName, projectName),
			},
			{
				ResourceName:      "rollbar_team_project_association.foobar",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccRollbarTeamProjectAssociationImportStateIdFunc(
					"rollbar_team_project_association.foobar"),
			},
		},
	})
}

func testAccRollbarTeamProjectAssociationImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("%s:%s", rs.Primary.Attributes["team_id"],
			rs.Primary.Attributes["project_id"]), nil
	}
}
