---
layout: "rollbar"
page_title: "Rollbar: rollbar_team_user_association"
sidebar_current: "docs-rollbar-resource-team-user-association"
description: |-
Provides a resource to create and manage the association between a team and user.
---

# rollbar_team_user_association

This resource is used to create and manage the association between a team and user.

### Unique resource lifecycle

If the specified `email` belongs to a new user:

* For resource creation, an invitation will be sent to the user to create a Rollbar account.
  Once the account is created, the user will join the team.
* For resource deletion, if the invitation has been accepted, and the user has joined the team,
  the user will be removed from the team. If the invitation has not been accepted, the invitation
  will be revoked.
* For resource state refresh, if the invitation was either cancelled or rejected, a new invitation
  will be sent out on the next resource creation unless the associated Terraform code is removed
  from your configuration.

If the specified `email` belongs to an existing Rollbar user:

* For resource creation, the user will be immediately added to the team.
* For resource deletion, the user will be removed from the team.

## Example Usage

```hcl-terraform
data "rollbar_team" "foobar" {
  id = "my_team_id"
}

data "rollbar_user" "foobar" {
  email = "user@email.com"
}

resource "rollbar_team_user_association" "foobar" {
  team_id = data.rollbar_team.foobar.id
  email   = data.rollbar_user.foobar.email
}
```

```hcl-terraform
resource "rollbar_team_user_association" "foobar" {
  team_id = 123456
  email   = "email_to_invite@company.com"
}
```

## Argument Reference

The following arguments are supported:

* `team_id` - (Required) `<string>` ID of existing team.
* `email` - (Required) `<string>` Email address of an existing Rollbar user or a new user.

## Attributes Reference

* `user_id` - Email address of a user.
* `invited_or_added` - Whether the user was either initially `invited` or `added`
  to the Rollbar team.
* `invitation_status` - Status of the invitation. This attribute is set only if the user
  had to first be invited to Rollbar in order to join the team.
* `invitation_id` - ID of the invitation. This attribute is set only if the user
  had to first be invited to Rollbar in order to join the team.

## Import

Existing team user association can be imported using a composite value of the team ID and email address
separated by a colon.

For example:

```shell
$ terraform import rollbar_team_user_association.follbar 123:user@company.com
```

-> **IMPORTANT!**
You can only import a team user association if the user has already joined the team.