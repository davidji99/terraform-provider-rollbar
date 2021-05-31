---
layout: "rollbar"
page_title: "Rollbar: rollbar_user"
sidebar_current: "docs-rollbar-datasource-user-x"
description: |-
  Get information on a Rollbar User.
---

# Data Source: rollbar_user

Use this data source to get information about a Rollbar user. The Rollbar user must be a member of the account
that is used to authenticate with the provider.

## Example Usage

```hcl-terraform
data "rollbar_user" "foobar" {
  email = "user@email.com"
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) The user email

## Attributes Reference

The following attributes are exported:

* `username` - The user's username
