---
layout: "rollbar"
page_title: "Rollbar: rollbar_project"
sidebar_current: "docs-rollbar-resource-project"
description: |-
  Provides a resource to create and manage a Rollbar project.
---

# rollbar\_project

This resource is used to create and manage projects on Rollbar.

~> NOTE: The Rollbar API does not support updating existing projects, only through the UI.
Therefore, you must update your configuration file(s) for this resource if you manually updated
the project name. Otherwise, your `terraform plan` will detect if a difference between the state file and remote.

## Example Usage

```hcl-terraform
# Create a new Rollbar project
resource "rollbar_project" "follbar" {
    name = "my_new_project"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) `<string>` Name of the project.

## Attributes Reference

The following attributes are exported:

* `status` - Whether the project is enabled or not. Returns a `string`, not `boolean`.
* `account_id` - The account the project belongs to.

## Import

Existing project(s) can be imported using the project id.

For example:

```
$ terraform import rollbar_project.follbar <PROJECT_ID>
```