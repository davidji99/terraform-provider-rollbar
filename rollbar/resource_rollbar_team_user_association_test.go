package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strings"
	"testing"
)

func TestAccRollbarTeamUserAssociation_BasicInvited(t *testing.T) {
	teamID := testAccConfig.GetTeamIDorAbort(t)

	emailSplitted := strings.Split(testAccConfig.GetTeamEmailAddress(t), "@")
	email := fmt.Sprintf("%s+%s@%s", emailSplitted[0], acctest.RandString(10), emailSplitted[1])

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarTeamUserAssociation_basic(teamID, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"rollbar_team_user_association.foobar", "team_id", teamID),
					resource.TestCheckResourceAttr(
						"rollbar_team_user_association.foobar", "email", email),
					//resource.TestCheckResourceAttrSet(
					//	"rollbar_team_user_association.foobar", "user_id"),
					resource.TestCheckResourceAttr(
						"rollbar_team_user_association.foobar", "invited_or_added", "invited"),
					resource.TestCheckResourceAttrSet(
						"rollbar_team_user_association.foobar", "invitation_status"),
					resource.TestCheckResourceAttrSet(
						"rollbar_team_user_association.foobar", "invitation_id"),
				),
			},
		},
	})
}

func TestAccRollbarTeamUserAssociation_BasicAdded(t *testing.T) {
	teamID := testAccConfig.GetTeamIDorAbort(t)
	email := testAccConfig.GetTeamEmailAddress(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarTeamUserAssociation_basic(teamID, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"rollbar_team_user_association.foobar", "team_id", teamID),
					resource.TestCheckResourceAttr(
						"rollbar_team_user_association.foobar", "email", email),
					resource.TestCheckResourceAttrSet(
						"rollbar_team_user_association.foobar", "user_id"),
					resource.TestCheckResourceAttr(
						"rollbar_team_user_association.foobar", "invited_or_added", "added"),
					//resource.TestCheckResourceAttrSet(
					//	"rollbar_team_user_association.foobar", "invitation_status"),
					//resource.TestCheckResourceAttrSet(
					//	"rollbar_team_user_association.foobar", "invitation_id"),
				),
			},
		},
	})
}

func testAccCheckRollbarTeamUserAssociation_basic(teamID, email string) string {
	return fmt.Sprintf(`
resource "rollbar_team_user_association" "foobar" {
	team_id = %s
	email = "%s"
}
`, teamID, email)
}
