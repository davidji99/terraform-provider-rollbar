---
layout: "rollbar"
page_title: "Rollbar: rollbar_team"
sidebar_current: "docs-rollbar-resource-team"
description: |-
  Provides a resource to create and manage a Rollbar team.
---

# rollbar\_team

This resource is used to create and manage teams on Rollbar.

**NOTE:** The Rollbar API does not support updating existing teams, only through the UI.
Therefore, you must update your configuration file(s) for this resource if you manually updated
the team name. Otherwise, your `terraform plan` will detect if a difference between the state file and remote.

## Example Usage

```hcl
# Create a new Rollbar team
resource "rollbar_team" "follbar" {
    name = "my_new_team"
    access_level = "standard"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) `<string>` Name of the team.

* `access_level` - (Required) `<string>` Access level of the team. Valid options: `standard`, `light`, `view`.
`standard` is the only access level you can choose in the UI. `light` and `view` are API-only team access levels.
`light` gives the team read and write access, but not to all settings. `view` gives the team read-only access.

## Attributes Reference

The following attributes are exported:

* `account_id` - The account the team belongs to.

## Import

Existing team(s) can be imported using the team id.

For example:

```
$ terraform import rollbar_team.follbar <TEAM_ID>
```