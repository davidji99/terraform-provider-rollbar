package rollbar

import (
	"context"
	"fmt"
	"github.com/davidji99/rollrest-go/rollrest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"regexp"
	"strconv"
)

const (
	TeamUserAddedStatus   = "added"
	TeamUserInvitedStatus = "invited"
)

func resourceRollbarTeamUserAssociation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRollbarTeamUserAssociationCreate,
		ReadContext:   resourceRollbarTeamUserAssociationRead,
		DeleteContext: resourceRollbarTeamUserAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceRollbarTeamUserAssociationImport,
		},

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"invited_or_added": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"invitation_status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"invitation_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceRollbarTeamUserAssociationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*Config).API

	result, parseErr := ParseCompositeID(d.Id(), 2)
	if parseErr != nil {
		return nil, parseErr
	}

	teamID, _ := strconv.Atoi(result[0])
	email := result[1]

	// Retrieve user ID by email
	user, _, userFindErr := findUserByEmail(client, email)
	if userFindErr != nil {
		return nil, fmt.Errorf("did not find an existing Rollbar user with email %s", email)
	}

	userID := int(user.GetID())

	// Check if user has been added to team
	isMember, _, err := client.Teams.IsUserMember(teamID, userID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, fmt.Errorf("cannot import - user %d has not been added to team %d", userID, teamID)
	}

	d.SetId(constructTeamUserResourceID(teamID, email, TeamUserAddedStatus))
	d.Set("email", email)
	d.Set("user_id", int(user.GetID()))
	d.Set("invited_or_added", TeamUserAddedStatus)
	d.Set("invitation_status", "")
	d.Set("invitation_id", 0)

	return []*schema.ResourceData{d}, nil
}

func findUserByEmail(client *rollrest.Client, email string) (*rollrest.User, bool, error) {
	users, _, userInfoErr := client.Users.List()
	if userInfoErr != nil {
		return nil, false, userInfoErr
	}

	for _, u := range users.GetResult().Users {
		if u.GetEmail() == email {
			return u, true, nil
		}
	}

	return nil, false, nil
}

func constructTeamUserResourceID(teamID int, email, status string) string {
	return fmt.Sprintf("%d:%s:%s", teamID, email, status)
}

func resourceRollbarTeamUserAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API
	teamID := getTeamID(d)
	email := getEmail(d)

	log.Printf("[DEBUG] Inviting or adding %s to team %d", email, teamID)

	inviteResponse, _, inviteErr := client.Teams.InviteUser(teamID, &rollrest.TeamInviteRequest{Email: email})
	if inviteErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("unable to invite/add %s to team %d", email, teamID),
			Detail:   inviteErr.Error(),
		})
		return diags
	}

	log.Printf("[DEBUG] Invited or added %s to team %d", email, teamID)

	var resourceID string
	var invitedOrAdded string

	// If the invite response returns the following message string, the email address belongs to an existing Rollbar user,
	// and that user will be immediately added to the team.
	if regexp.MustCompile(`given email address has been added`).MatchString(inviteResponse.GetMessage()) {
		resourceID = constructTeamUserResourceID(teamID, email, TeamUserAddedStatus)
		invitedOrAdded = TeamUserAddedStatus
	}

	// If the invite response returns an invitation, the email address has been sent an invitation and the user needs
	// to accept it before being added to the team.
	if inviteResponse.GetResult() != nil {
		resourceID = constructTeamUserResourceID(teamID, email, TeamUserInvitedStatus)
		invitedOrAdded = TeamUserInvitedStatus
	}

	if resourceID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary: fmt.Sprintf("Invite/added %s to team %d but the API response is not expected. Likely a provider issue.",
				email, teamID),
			Detail: fmt.Sprintf("response error count: %d | message: %s | result: %v",
				inviteResponse.GetErrorCount(), inviteResponse.GetMessage(), inviteResponse.GetResult()),
		})
		return diags
	}

	d.SetId(resourceID)

	// This will either be the actual invitation ID or an empty string
	d.Set("invitation_id", int(inviteResponse.GetResult().GetID()))
	d.Set("invited_or_added", invitedOrAdded)

	return resourceRollbarTeamUserAssociationRead(ctx, d, meta)
}

func resourceRollbarTeamUserAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API

	result, parseErr := ParseCompositeID(d.Id(), 3)
	if parseErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to parse resource ID during state refresh",
			Detail:   parseErr.Error(),
		})
		return diags
	}

	teamID, _ := strconv.Atoi(result[0])
	email := result[1]
	status := result[2]

	d.Set("team_id", teamID)
	d.Set("email", email)
	d.Set("invitation_status", "")

	if status == TeamUserAddedStatus || d.Get("invitation_status").(string) == "accepted" {
		user, _, userFindErr := findUserByEmail(client, email)
		if userFindErr != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("cannot determine if %s exists in Rollbar", email),
				Detail:   userFindErr.Error(),
			})
			return diags
		}
		d.Set("user_id", int(user.GetID()))
	}

	if status == TeamUserInvitedStatus {
		inviteID := d.Get("invitation_id").(int)
		inviteStatus, _, statusErr := client.Invitations.Get(inviteID)
		if statusErr != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("unable to retrieve invitation %d during state refresh", inviteID),
				Detail:   statusErr.Error(),
			})
			return diags
		}
		d.Set("invitation_status", inviteStatus.GetResult().GetStatus())
	}

	return diags
}

func resourceRollbarTeamUserAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API

	result, parseErr := ParseCompositeID(d.Id(), 3)
	if parseErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to parse resource ID during state refresh",
			Detail:   parseErr.Error(),
		})
		return diags
	}

	teamID, _ := strconv.Atoi(result[0])
	email := result[1]
	status := result[2]

	if status == TeamUserInvitedStatus {
		inviteID := d.Get("invitation_id").(int)

		log.Printf("[DEBUG] Cancelling invitation %d", inviteID)

		// Cancel the invitation
		_, _, cancelErr := client.Invitations.Cancel(d.Get("invitation_id").(int))
		if cancelErr != nil {
			log.Printf("[DEBUG] issue cancelling invitation but that's okay as subsequent invitations to the same email address will invalidate any pending ones")
		}

		log.Printf("[DEBUG] Cancelled invitation %d", inviteID)
	}

	if status == TeamUserAddedStatus {
		log.Printf("[DEBUG] Removing %s from team %d", email, teamID)

		_, _, removeErr := client.Teams.RemoveUser(teamID, d.Get("user_id").(int))
		if removeErr != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("could not remove %s from team %d", email, teamID),
				Detail:   removeErr.Error(),
			})
			return diags
		}

		log.Printf("[DEBUG] Removed %s from team %d", email, teamID)
	}

	d.SetId("")

	return diags
}
