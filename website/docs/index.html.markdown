---
layout: "rollbar"
page_title: "Provider: Rollbar"
sidebar_current: "docs-rollbar-index"
description: |-
  The Rollbar provider is used to interact with the resources provided by the Rollbar API.
---

# Rollbar Provider

The Rollbar provider is used to interact with the resources provided by [Rollbar API](https://docs.rollbar.com/reference).
and needs to be configured with credentials before it can be used. This provider has been developed
using the [terraform sdk](https://github.com/hashicorp/terraform-plugin-sdk) and is recommended to be used with `terraform v0.12.X+`.

## Background

[Rollbar](https://rollbar.com) automates error monitoring and triaging, so developers can fix errors that matter within minutes,
and build software quickly and painlessly.

## Contributing

Development happens in the [GitHub repo](https://github.com/davidji99/terraform-provider-rollbar):

* [Releases](https://github.com/davidji99/terraform-provider-rollbar/releases)
* [Changelog](https://github.com/davidji99/terraform-provider-rollbar/blob/master/CHANGELOG.md)
* [Issues](https://github.com/davidji99/terraform-provider-rollbar/issues)

## Example Usage

```hcl
# Configure the Rollbar provider
provider "rollbar" {
  account_access_token = "some_token"
}

# Create a new project
resource "rollbar_project" "service-x" {
  # ...
}
```

## Authentication

Certain resources in Rollbar require either the `account_access_token` or `project_access_token`. Based on observation,
the `account_access_token` is used more frequently for this provider's resources. You must supply both tokens 
if your terraform configuration code manages resources that require both access tokens. Otherwise, one access token
must be supplied to your provider block or sourced from other means.

The Rollbar provider offers a flexible means of providing credentials for authentication.
The following methods are supported, listed in order of precedence, and explained below:

* Static credentials
* Environment variables

### Static credentials

Credentials can be provided statically by adding `account_access_token` argument
to the Rollbar provider block:

```hcl
provider "rollbar" {
  account_access_token = "some_token"
}
```

### Environment variables

When the Rollbar provider block does not contain an `account_access_token`
argument, the missing credentials will be sourced from the environment via the
`ROLLBAR_ACCOUNT_ACCESS_TOKEN` environment variables:

```hcl
provider "rollbar" {}
```

```shell
$ export ROLLBAR_ACCOUNT_ACCESS_TOKEN="some_token"
$ terraform plan
Refreshing Terraform state in-memory prior to plan...
```

## Argument Reference

The following arguments are supported:

* `account_access_token` - (Required) Rollbar account access token. It can be provided, but it can also
be sourced from [other locations](#Authentication). This token **MUST** have read & write permissions enabled
so the provider can completely manage supported resources.

* `project_access_token` - (Required) Rollbar project access token. It can be provided, but it can also
be sourced from [other locations](#Authentication). This token **MUST** have read & write permissions enabled
so the provider can completely manage supported resources.

* `api_headers` - (Optional) Additional API headers.

* `post_create_pd_integration_delete_default_rules` - (Optional) Delete the auto-added rules after enabling
PagerDuty notification integration. Defaults to `false`.