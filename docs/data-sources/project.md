---
layout: "rollbar"
page_title: "Rollbar: rollbar_project"
sidebar_current: "docs-rollbar-datasource-project-x"
description: |-
  Get information on a Rollbar Project.
---

# Data Source: rollbar_project

Use this data source to get information about a Rollbar Project.

## Example Usage

```hcl-terraform
data "rollbar_project" "foobar" {
  name = "my_project"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The project name

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The project id
* `status` - The project status
* `account_id` - The account id the project belongs to
