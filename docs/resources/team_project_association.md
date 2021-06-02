---
layout: "rollbar"
page_title: "Rollbar: rollbar_team_project_association"
sidebar_current: "docs-rollbar-resource-team-project-association"
description: |-
Provides a resource to create and manage the association between a team and project.
---

# rollbar_team_project_association

This resource is used to create and manage the association between a team and project.

## Example Usage

```hcl-terraform
resource "rollbar_team" "foobar" {
  name = "my_team"
  access_level = "standard"
}

resource "rollbar_project" "foobar" {
  name = "my_project"
}

resource "rollbar_team_project_association" "foobar" {
  team_id = rollbar_team.foobar.id
  project_id = rollbar_project.foobar.id
}
```

## Argument Reference

The following arguments are supported:

* `team_id` - (Required) `<string>` ID of existing team.
* `project_id` - (Required) `<string>` ID of existing project.

## Attributes Reference

n/a

## Import

Existing team project association can be imported using a composite value of the team and project ID
separated by a colon.

For example:

```shell
$ terraform import rollbar_team_project_association.follbar 123:456
```