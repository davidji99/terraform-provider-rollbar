---
layout: "rollbar"
page_title: "Rollbar: rollbar_pagerduty_notification_rule"
sidebar_current: "docs-rollbar-resource-pagerduty-notification-rule"
description: |-
  Provides a resource to create and partially manage Rollbar PagerDuty notification rules.
---

# rollbar\_pagerduty\_notification\_rule

This resource is used to manage Rollbar's PagerDuty notification rules. You must supply a `project_access_token` with write
permissions in other to manage this resource.

For more information on the supported values when constructing a rule, please visit [this page](https://docs.rollbar.com/reference#setup-pagerduty-notification-rules).
(Yes, you'll have to read an OpenAPI spec.)

~> NOTE: Due to API limitations, it is not possible to selectively `GET`, `DELETE` or `CREATE` a single notification rule.
Whatever rule(s) you define in your terraform configuration **will be the only rules** present in your account 
after a `terraform apply`. This is especially important to understand if you have pre-existing rules in your account 
prior to terraform managing this resource or rules created outside of terraform. In other words, this provider/terraform 
will overwrite any remotely defined rules not in your configuration files. Furthermore, it is strongly advised that you only declare one `rollbar_pagerduty_notification_rule` in your configuration files for the reasons above.

## Example Usage

```hcl
# Create a new Rollbar PagerDuty notification rule
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
		config = {
			service_key = "aG59dD4FtWRfGMNJ3mLcZTK3CC4Qhgas" // this is not a real `service_key`
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `rule` - (Required) A PagerDuty notification rule

    * `trigger` - (Required) `<string>` Valid options are: `new_item`, `occurrence_rate`, `resolved_item`,
    `reactivated_item`, `exp_repeat_item`.

    * `filter` - (Required)

        * `type` - (Required) `<string>` The type of rule filter. 
        Valid options are: `environment`, `level`, `title`, `filename`,`context`, `method`, `framework`, `path`,
        `rate`,`unique_occurrences`.

        * `operation` - (Required) `<string>`

        * `value` - (Optional) `<string>` This attribute is **required** for all filter types except for when constructing
        a 'path_filter_with_exists' filter. See the API documentation (OpenAPI Spec) for more information. The attribute
        itself is marked as 'Optional' to support a 'path_filter_with_exists' filter.

        * `path` - (Optional) `<string>` This attribute is used only for filter type `path`.

        * `period` - (Optional) `<integer>` Number of seconds.

        * `count` - (Optional) `<integer>` Rate of occurrences of an item.

    * `config` - (Optional) Any additional rule configurations

        * `service_key` - (Required) `string` Use this service API key instead of the default PagerDuty Service API key.

#### Rule Explanation
Certain rule triggers will require certain filters. Here are some of the following requirements:
1. Define a `filter.type` of `rate` when constructing an `occurrence_rate` rule trigger.

Please refer to the official [Rollbar OpenAPI specification](https://explorer.docs.rollbar.com/main.yaml) for more information.

#### Filter Options

As of Feb 11th, 2020, the Rollbar Rest API documentation does not present the available options to the `rule.filter` 
attribute block in an easily readable manner. Therefore, this section will provide a summary of the available 
options for `rule.filter`:

1. For `filter.type` of `environment`:
    * Valid `filter.operation` option(s): `eq`, `neq`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` of `level`:
       * Valid `filter.operation` option(s): `eq`, `gte`
       * Valid `filter.value` option(s): `debug`, `info`, `warning`, `error`, `critical` (case sensitive!)

1. For `filter.type` of `title`:
      * Valid `filter.operation` option(s): `within`, `nwithin`, `regex`, `nregex`
      * Valid `filter.value` option(s): any freeform `string`*
      
1. For `filter.type` of `filename`:
    * Valid `filter.operation` option(s): `within`, `nwithin`, `regex`, `nregex`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` of `context`:
    * Valid `filter.operation` option(s): `startswith`, `eq`, `neq`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` of `method`:
    * Valid `filter.operation` option(s): `within`, `nwithin`, `regex`, `nregex`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` of `framework`:
    * Valid `filter.operation` option(s): `eq`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` of `rate`:
    * Valid `filter.period` option(s): Whole number greater than zero
    * Valid `filter.count` option(s): Whole number greater than zero

1. For `filter.type` of `unique_occurrences`:
    * Valid `filter.operation` option(s): Only valid option is `gte`
    * Valid `filter.value` option(s): Whole number greater than zero

For the `filter.type` of `path`, there are two possible setups:

* If you define `filter.value` & `filter.operation` in your terraform configuration,
    * Valid `filter.operation` option(s): `eq`, `gte`, `lte`, `within`, `nwithin`, `neq`, `regex`, `nregex`, `startswith`
    * Valid `filter.value` option(s): any freeform `string`*
    * Valid `filter.path` option(s): any freeform `string`*

* If you only define `filter.operation` in your terraform configuration _(aka 'path_filter_with_exists)_,
    * Valid `filter.operation` option(s): `eq`, `gte`, `lte`, `within`, `nwithin`, `neq`, `regex`, `nregex`, `startswith`, 
    `exists`, `nexists`
    * Valid `filter.value` option(s): any freeform `string`* & is Optional.
    * Valid `filter.path` option(s): any freeform `string`*

**Note:** Any `filter.value` with an asterisk may not be fully accurate. Please open an issue on the GitHub repository
if you encounter any problems.

## Attributes Reference

N/A

## Import

Due to API limitations, it is not possible to import an existing PagerDuty notification rule(s).
