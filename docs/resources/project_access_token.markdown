---
layout: "rollbar"
page_title: "Rollbar: rollbar_project_access_token"
sidebar_current: "docs-rollbar-resource-project-access-token"
description: |-
  Provides a resource to create and manage a Rollbar project access token.
---

# rollbar\_project\_access\_token

This resource is used to create and manage access tokens for a Rollbar project.

Access tokens are used to read and write data via the Rollbar API.
Each token has a set of scopes that determine which operations can be performed with the token.

Tokens have configurable rate limits that define how many times they may be used for API calls in a given interval.
The default rate limit and system max is 5000 calls every 1 minute. The default rate limit takes precedence over any custom rate limits.
If you need a higher rate limit, please contact support@rollbar.com.

**NOTE:** As of (August 6th, 2020), the Rollbar API does not provide support for the following
and therefore cannot be implemented in the provider:

1. Deleting access tokens. Instead, the provider will 'disable' the token by setting its rate limit to 1 call per 30 days
and remove the resource from your state. Then, the user can delete the token in the Rollbar UI.
1. Updating a project access token's `name`, `status`, and `scopes`. Users will need to make updates via the UI
and then update their terraform configuration prior to a `plan` or `apply`. Otherwise, terraform will detect a diff
that cannot be resolved by any terraform `apply`.

Please also note that a project, by default, comes with four project access tokens each only have one of the four scopes. If you wish to use those tokens instead of creating new ones, it is recommended to use the `rollbar_project_access_tokens` data source.

## Example Usage

```hcl
# Create a new Rollbar project access token
resource "rollbar_project" "foobar" {
	name = "new_project"
}

resource "rollbar_project_access_token" "foobar" {
	project_id = rollbar_project.foobar.id
	name = "read only"
	scopes = ["read"]
	status = "enabled"
	rate_limit_window_size = 60
	rate_limit_window_count = 1500
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) `<string>` The ID of the project

* `name` - (Required) `<string>` Name of the project access token. Max length 32 characters.

* `scopes` - (Required) `<list(string)>` Scopes to assign to the create access token.
Valid options: `read`, `write`, `post_server_item`, `post_client_server`.

* `status` - (Required) `<string>` Enable or disable the access token. Valid options: `enabled`, `disabled`.

* `rate_limit_window_size` `<integer>` - Period of time (in seconds) for the rate limit. On **resource creation only**,
the valid options are the following: `0, 60, 300, 1800, 3600, 86400, 604800, 2592000`.
Otherwise, any value greater than `0`. If this argument is not set, the default is 60 seconds (1 minute).

* `rate_limit_window_count` `<integer>` - Number of requests for the defined rate limiting period.
Otherwise, any value greater than `0`. If this argument is not set, the default is 5000 calls.

## Attributes Reference

The following attributes are exported:

* `cur_rate_limit_window_count` - How many remaining API calls are left for the access token.

* `date_created` - The timestamp in epoch of when the token was created.

* `access_token` - The actual access token. This value is set to `Sensitive`
and will not be shown in any non-debug `terraform` outputs.

## Import

Existing project access tokens(s) can be imported using a combination of the project id & access token separated by a colon.

For example:

```
$ terraform import rollbar_project_access_token.follbar <PROJECT_ID>:<ACCESS_TOKEN>
```

The provider will then lookup the access token and set a randomly generated number for the resource ID during importation.
At no point in this resource's lifecycle will the ID be set to the access token to avoid plaintexting a secret.
