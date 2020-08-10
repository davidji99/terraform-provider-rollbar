package rollbar

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
)

func TestAccRollbarPagerDutyNotificationRule_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRollbarPagerDutyNotificationRule_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"rollbar_pagerduty_notification_rule.foobar", "rule.0.filter.0.operation", "gte"),
					resource.TestCheckResourceAttr(
						"rollbar_pagerduty_notification_rule.foobar", "rule.1.config.0.service_key", "aG59dD4FtWRfGMNJ3mLcZTK3CC4Qhgas"),
				),
			},
		},
	})
}

func TestAccRollbarPagerDutyNotificationRule_InvalidTrigger(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckRollbarPagerDutyNotificationRule_InvalidTrigger(),
				ExpectError: regexp.MustCompile(`expected rule.0.trigger to be one of \[new_item occurrence_rate resolved_item reactivated_item exp_repeat_item\], got new_item123`),
			},
		},
	})
}

func testAccCheckRollbarPagerDutyNotificationRule_InvalidTrigger() string {
	return fmt.Sprintf(`
resource "rollbar_pagerduty_notification_rule" "foobar" {
	rule {
		trigger = "new_item123"
		filter {
			type = "level"
			operation = "gte"
			value = "critical"
		}
	}
}
`)
}

func testAccCheckRollbarPagerDutyNotificationRule_basic() string {
	return fmt.Sprintf(`
resource "rollbar_pagerduty_notification_rule" "foobar" {
	rule {
		trigger = "new_item"
		filter {
			type = "level"
			operation = "gte"
			value = "critical"
		}
		filter {
			type = "title"
			operation = "within"
			value = "some_title"
		}
	}

	rule {
		trigger = "new_item"
		filter {
			type = "environment"
			operation = "eq"
			value = "production"
		}
		config {
			service_key = "aG59dD4FtWRfGMNJ3mLcZTK3CC4Qhgas"
		}
	}
}
`)
}
