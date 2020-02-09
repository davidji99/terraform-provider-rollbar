---
layout: "rollbar"
page_title: "Rollbar: rollbar_team"
sidebar_current: "docs-rollbar-datasource-team-x"
description: |-
  Get information on a Rollbar Team.
---

# Data Source: rollbar_team

Use this data source to get information about a Rollbar Team.

## Example Usage

```hcl
data "rollbar_team" "foobar" {
  id = "<SOME_TEAM_ID>"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The team id

## Attributes Reference

The following attributes are exported:

* `name` - The team name
* `access_level` - The access level
