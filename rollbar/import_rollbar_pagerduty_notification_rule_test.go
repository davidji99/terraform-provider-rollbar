package rollbar

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
)

func TestAccRollbarPagerDutyNotificationRule_importBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarPagerDutyNotificationRule_basic(),
			},
			{
				ResourceName:      "rollbar_pagerduty_notification_rule.foobar",
				ExpectError:       regexp.MustCompile(`not possible to import PagerDuty`),
				ImportStateVerify: true,
				ImportState:       true,
			},
		},
	})
}
