package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccRollbarTeamProjectAssociation_Basic(t *testing.T) {
	teamName := fmt.Sprintf("team-%s", acctest.RandString(10))
	projectName := fmt.Sprintf("team-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarTeamProjectAssociation_basic(teamName, projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"rollbar_team_project_association.foobar", "team_id"),
					resource.TestCheckResourceAttrSet(
						"rollbar_team_project_association.foobar", "project_id"),
				),
			},
		},
	})
}

func testAccCheckRollbarTeamProjectAssociation_basic(teamName, projectName string) string {
	return fmt.Sprintf(`
resource "rollbar_team" "foobar" {
	name = "%s"
	access_level = "standard"
}

resource "rollbar_project" "foobar" {
	name = "%s"
}

resource "rollbar_team_project_association" "foobar" {
	team_id = rollbar_team.foobar.id
	project_id = rollbar_project.foobar.id
}
`, teamName, projectName)
}
