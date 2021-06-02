---
layout: "rollbar"
page_title: "Rollbar: rollbar_pagerduty_integration"
sidebar_current: "docs-rollbar-resource-pagerduty-integration"
description: |-
  Provides a resource to create and partially manage a Rollbar PagerDuty integration.
---

# rollbar\_pagerduty\_integration

This resource is used to manage Rollbar's integration with PagerDuty. You must supply a `project_access_token` with write
permissions in other to manage this resource.

~> NOTE: Due to API limitations, it is not possible to delete/remove the integration via the API.
Therefore upon resource deletion, the existing PagerDuty integration will be disabled. Users must then visit the UI
if they wish to remove the integration entirely by clearing out the 'Service API Key' field and click 'Save'.

Users also have the option to set `post_create_pd_integration_delete_default_rules` to `true` in their `provider` block
if they wish to delete the auto-added notification rules. This is recommended as the `rollbar_pagerduty_notification_rule`
resource cannot import existing rules due to API limitations.

## Example Usage

```hcl-terraform
# Create a new Rollbar PagerDuty Integration
resource "rollbar_pagerduty_integration" "pd" {
	service_key = "SOME_VALID_PD_KEY"
	enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `service_key` - (Required) `<string>` Valid PagerDuty Service API Key. Must 32 characters long.
* `enabled` - (Required) `<boolean>` Enable the PagerDuty notifications globally

## Attributes Reference

N/A

## Import

Due to API limitations, it is not possible to import an existing PagerDuty integration.