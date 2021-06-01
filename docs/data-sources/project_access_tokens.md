---
layout: "rollbar"
page_title: "Rollbar: rollbar_project_access_tokens"
sidebar_current: "docs-rollbar-datasource-project-access-tokens-x"
description: |-
  Get information on all of a Rollbar project's access tokens.
---

# Data Source: rollbar_project_access_tokens

Use this data source to get all of a project's active access tokens.

If you wish only to use the a project's four default access tokens,
it is recommended to use data source to reference them in your terraform configuration.

## Example Usage

```hcl-terraform
resource "rollbar_project" "foobar" {
	name = "some_project"
}

data "rollbar_project_access_tokens" "foobar" {
  project_id = rollbar_project.foobar.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The project id

## Attributes Reference

The following attributes are exported:

* `access_tokens` - A map of access tokens where the key is the token name and the value is the access token.
