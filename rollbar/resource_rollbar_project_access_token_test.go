package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"regexp"
	"testing"
)

func TestAccRollbarProjectAccessToken_Basic(t *testing.T) {
	projectName := fmt.Sprintf("project-tftest-%s", acctest.RandString(10))
	tokenName := fmt.Sprintf("token-tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarProjectAccessToken_basic(projectName, tokenName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"rollbar_project_access_token.foobar", "name", tokenName),
					resource.TestCheckResourceAttr(
						"rollbar_project_access_token.foobar", "rate_limit_window_size", "60"),
					resource.TestCheckResourceAttr(
						"rollbar_project_access_token.foobar", "rate_limit_window_count", "1500"),
				),
			},
			{
				Config: testAccCheckRollbarProjectAccessToken_NewWinSize(projectName, tokenName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"rollbar_project_access_token.foobar", "rate_limit_window_size", "131"),
				),
			},
		},
	})
}

func TestAccRollbarProjectAccessToken_InvalidScopes(t *testing.T) {
	projectName := fmt.Sprintf("project-tftest-%s", acctest.RandString(10))
	tokenName := fmt.Sprintf("token-tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckRollbarProjectAccessToken_InvalidScopes(projectName, tokenName),
				ExpectError: regexp.MustCompile(`.*to be one of \[read write post_server_item post_client_server].*`),
			},
		},
	})
}

func TestAccRollbarProjectAccessToken_InvalidWinSizeOnCreation(t *testing.T) {
	projectName := fmt.Sprintf("project-tftest-%s", acctest.RandString(10))
	tokenName := fmt.Sprintf("token-tftest-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckRollbarProjectAccessToken_InvalidWinSize(projectName, tokenName),
				ExpectError: regexp.MustCompile(`.*is not a supported window size for token creation. Valid values are: \[0 60 300 1800 3600 86400 604800 2592000].*`),
			},
		},
	})
}

func testAccCheckRollbarProjectAccessToken_basic(projectName, tokenName string) string {
	return fmt.Sprintf(`
resource "rollbar_project" "foobar" {
	name = "%s"
}

resource "rollbar_project_access_token" "foobar" {
	project_id = rollbar_project.foobar.id
	name = "%s"
	scopes = ["read"]
	status = "enabled"
	rate_limit_window_size = 60
	rate_limit_window_count = 1500
}
`, projectName, tokenName)
}

func testAccCheckRollbarProjectAccessToken_InvalidScopes(projectName, tokenName string) string {
	return fmt.Sprintf(`
resource "rollbar_project" "foobar" {
	name = "%s"
}

resource "rollbar_project_access_token" "foobar" {
	project_id = rollbar_project.foobar.id
	name = "%s"
	scopes = ["read123", "dadad"]
	status = "enabled"
	rate_limit_window_size = 60
	rate_limit_window_count = 1500
}
`, projectName, tokenName)
}

func testAccCheckRollbarProjectAccessToken_InvalidWinSize(projectName, tokenName string) string {
	return fmt.Sprintf(`
resource "rollbar_project" "foobar" {
	name = "%s"
}

resource "rollbar_project_access_token" "foobar" {
	project_id = rollbar_project.foobar.id
	name = "%s"
	scopes = ["read"]
	status = "enabled"
	rate_limit_window_size = 59
	rate_limit_window_count = 1500
}
`, projectName, tokenName)
}

func testAccCheckRollbarProjectAccessToken_NewWinSize(projectName, tokenName string) string {
	return fmt.Sprintf(`
resource "rollbar_project" "foobar" {
	name = "%s"
}

resource "rollbar_project_access_token" "foobar" {
	project_id = rollbar_project.foobar.id
	name = "%s"
	scopes = ["read"]
	status = "enabled"
	rate_limit_window_size = 131
	rate_limit_window_count = 1500
}
`, projectName, tokenName)
}
