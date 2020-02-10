---
layout: "rollbar"
page_title: "Rollbar: rollbar_user"
sidebar_current: "docs-rollbar-datasource-user-x"
description: |-
  Get information on a Rollbar User.
---

# Data Source: rollbar_user

Use this data source to get information about a Rollbar user.

## Example Usage

```hcl
data "rollbar_user" "foobar" {
  id = "<SOME_USER_ID>"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The user id

## Attributes Reference

The following attributes are exported:

* `email` - The user's email address
* `username` - The user's username
