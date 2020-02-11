---
layout: "rollbar"
page_title: "Rollbar: rollbar_project_access_token"
sidebar_current: "docs-rollbar-datasource-project-access-token-x"
description: |-
  Get information on a Rollbar project's access token(s).
---

# Data Source: rollbar_project_access_token

Use this data source to get all of a project's active access token.

## Example Usage

```hcl
resource "rollbar_project" "foobar" {
	name = "some_project"
}

data "rollbar_project_access_token" "foobar" {
  project_id = rollbar_project.foobar.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The project id

## Attributes Reference

The following attributes are exported:

* `access_tokens` - A map of access tokens where the key is the token name and the value is the token.
