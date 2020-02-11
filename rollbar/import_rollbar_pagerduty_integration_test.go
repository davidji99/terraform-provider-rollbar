package rollbar

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"regexp"
	"testing"
)

func TestAccRollbarPagerDutyIntegration_importBasic(t *testing.T) {
	key := testAccConfig.GetPagerDutyAPIKeyorAbort(t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarPagerDutyIntegration_basic(key),
			},
			{
				ResourceName:      "rollbar_pagerduty_integration.foobar",
				ExpectError:       regexp.MustCompile(`not possible to import PagerDuty integration due to API limitations`),
				ImportStateVerify: true,
				ImportState:       true,
			},
		},
	})
}
