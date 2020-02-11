package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"regexp"
	"testing"
)

func TestAccRollbarPagerDutyIntegration_Basic(t *testing.T) {
	key := testAccConfig.GetPagerDutyAPIKeyorAbort(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarPagerDutyIntegration_basic(key),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"rollbar_pagerduty_integration.foobar", "service_key", key),
					resource.TestCheckResourceAttr(
						"rollbar_pagerduty_integration.foobar", "enabled", "true"),
				),
			},
		},
	})
}

func TestAccRollbarPagerDutyIntegration_InvalidServiceKey(t *testing.T) {
	key := "invalid_key"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckRollbarPagerDutyIntegration_basic(key),
				ExpectError: regexp.MustCompile(`expected length of service_key to be in the range \(32 - 32\)`),
			},
		},
	})
}

func testAccCheckRollbarPagerDutyIntegration_basic(key string) string {
	return fmt.Sprintf(`
resource "rollbar_pagerduty_integration" "foobar" {
	service_key = "%s"
	enabled = true
}
`, key)
}
