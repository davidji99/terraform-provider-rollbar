package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccDatasourceRollbarProject_Basic(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarProjectWithDatasourceBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.rollbar_project.foobar", "id"),
					resource.TestCheckResourceAttrSet(
						"data.rollbar_project.foobar", "name"),
					resource.TestCheckResourceAttrSet(
						"data.rollbar_project.foobar", "status"),
					resource.TestCheckResourceAttrSet(
						"data.rollbar_project.foobar", "account_id"),
				),
			},
		},
	})
}

func testAccCheckRollbarProjectWithDatasourceBasic(projectName string) string {
	return fmt.Sprintf(`
resource "rollbar_project" "foobar" {
	name = "%s"
}

data "rollbar_project" "foobar" {
  name = rollbar_project.foobar.name
}
`, projectName)
}
