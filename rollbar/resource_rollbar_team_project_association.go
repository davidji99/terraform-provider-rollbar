package rollbar

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceRollbarTeamProjectAssociation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRollbarTeamProjectAssociationCreate,
		ReadContext:   resourceRollbarTeamProjectAssociationRead,
		DeleteContext: resourceRollbarTeamProjectAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceRollbarTeamProjectAssociationImport,
		},

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRollbarTeamProjectAssociationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*Config).API

	result, _ := ParseCompositeID(d.Id(), 2)
	teamID, _ := strconv.Atoi(result[0])
	projectID, _ := strconv.Atoi(result[1])

	hasProject, _, err := client.Teams.HasProject(teamID, projectID)
	if err != nil {
		return nil, err
	}

	if !hasProject {
		return nil, fmt.Errorf("could not find project %d on team %d", projectID, teamID)
	}

	d.SetId(fmt.Sprintf("%s:%s", result[0], result[1]))
	d.Set("team_id", teamID)
	d.Set("project_id", projectID)

	return []*schema.ResourceData{d}, nil
}

func resourceRollbarTeamProjectAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API
	teamID := getTeamID(d)
	projectID := getProjectID(d)

	result, _, assignErr := client.Teams.AssignProject(teamID, projectID)
	if assignErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to associate project to team",
			Detail:   assignErr.Error(),
		})
		return diags
	}

	d.SetId(fmt.Sprintf("%d:%d", int(result.GetResult().GetTeamID()), int(result.GetResult().GetProjectID())))

	return resourceRollbarTeamProjectAssociationRead(ctx, d, meta)
}

func resourceRollbarTeamProjectAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API

	teamID := getTeamID(d)
	projectID := getProjectID(d)

	hasProject, _, err := client.Teams.HasProject(teamID, projectID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to verify team project association",
			Detail:   err.Error(),
		})
		return diags
	}

	if !hasProject {
		return diag.Errorf("could not find project %d on team %d", projectID, teamID)
	}

	return diags
}

func resourceRollbarTeamProjectAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*Config).API

	result, _ := ParseCompositeID(d.Id(), 2)
	teamID, _ := strconv.Atoi(result[0])
	projectID, _ := strconv.Atoi(result[1])

	_, removeErr := client.Teams.RemoveProject(teamID, projectID)
	if removeErr != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to disassociate project from team",
			Detail:   removeErr.Error(),
		})
		return diags
	}

	d.SetId("")

	return diags
}
