package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"regexp"
	"testing"
)

func TestAccRollbarProject_Basic(t *testing.T) {
	name := fmt.Sprintf("tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarProject_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"rollbar_project.foobar", "name", name),
				),
			},
		},
	})
}

func TestAccRollbarProject_InvalidName(t *testing.T) {
	name := "invalid projectname@"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarProject_basic(name),
				ExpectError: regexp.MustCompile("Must start with a letter and can only contain letters, numbers, underscores, " +
					"hyphens, and periods. Max length 32 characters."),
			},
		},
	})
}

func testAccCheckRollbarProject_basic(name string) string {
	return fmt.Sprintf(`
resource "rollbar_project" "foobar" {
	name = "%s"
}
`, name)
}
