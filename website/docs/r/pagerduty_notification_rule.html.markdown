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
will overwrite any remotely defined rules not in your configuration files.

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

    * `trigger` - (Required) `<string>` The only valid option is `new_item` for now.

    * `filter` - (Required)

        * `type` - (Required) `<string>` The type of rule filter. 
        Valid options are: `environment`, `level`, `title`, `filename`,`context`, `method`, `framework`, `path`.

        * `operation` - (Required) `<string>`

        * `value` - (Optional) `<string>` This attribute is **required** for all filter types except for when constructing
        a 'path_filter_with_exists' filter. See the API documentation (OpenAPI Spec) for more information. The attribute
        itself is marked as 'Optional' to support a 'path_filter_with_exists' filter.

        * `path` - (Optional) `<string>` This attribute is used only for filter type `path`.

    * `config` - (Optional) Any additional rule configurations

        * `service_key` - (Required) `string` Use this service API key instead of the default PagerDuty Service API key.

### Filter Explanations

As of Feb 11th, 2020, the Rollbar Rest API documentation does not present the available options to the `rule.filter` 
attribute block in an easily readable manner. Therefore, this section will provide a summary of the available 
options for `rule.filter`:

1. For `filter.type` type `environment`:
    * Valid `filter.operation` option(s): `eq`, `neq`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` type `level`:
       * Valid `filter.operation` option(s): `eq`, `gte`
       * Valid `filter.value` option(s): `debug`, `info`, `warning`, `error`, `critical` (case sensitive!)

1. For `filter.type` type `title`:
      * Valid `filter.operation` option(s): `within`, `nwithin`, `regex`, `nregex`
      * Valid `filter.value` option(s): any freeform `string`*
      
1. For `filter.type` type `filename`:
    * Valid `filter.operation` option(s): `within`, `nwithin`, `regex`, `nregex`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` type `context`:
    * Valid `filter.operation` option(s): `startswith`, `eq`, `neq`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` type `method`:
    * Valid `filter.operation` option(s): `within`, `nwithin`, `regex`, `nregex`
    * Valid `filter.value` option(s): any freeform `string`*

1. For `filter.type` type `framework`:
    * Valid `filter.operation` option(s): `eq`
    * Valid `filter.value` option(s): any freeform `string`*

For the `filter.type` type `path`, there are two possible setups:

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