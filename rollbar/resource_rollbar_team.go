package rollbar

import (
	"github.com/davidji99/rollrest-go/rollrest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"regexp"
)

func resourceRollbarTeam() *schema.Resource {
	return &schema.Resource{
		Create: resourceRollbarTeamCreate,
		Read:   resourceRollbarTeamRead,
		Delete: resourceRollbarTeamDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRollbarTeamImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z][0-9A-Za-z,.\-_]{1,31}$`),
					"Must start with a letter and can only contain letters, numbers, underscores, "+
						"hyphens, and periods. Max length 32 characters."),
			},

			"access_level": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"standard", "light", "view"}, false),
			},

			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceRollbarTeamImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	readErr := resourceRollbarTeamRead(d, meta)

	return []*schema.ResourceData{d}, readErr
}

func resourceRollbarTeamCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API
	opts := &rollrest.TeamRequest{}

	if v, ok := d.GetOk("name"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] team name is : %s", vs)
		opts.Name = vs
	}

	if v, ok := d.GetOk("access_level"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] team name is : %s", vs)
		opts.AccessLevel = vs
	}

	log.Printf("[DEBUG] Creating new team %s", opts.Name)

	newTeam, _, createErr := client.Teams.Create(opts)
	if createErr != nil {
		return createErr
	}

	log.Printf("[DEBUG] Created new team %s", opts.Name)

	d.SetId(Int64ToString(newTeam.GetResult().GetID()))

	return resourceRollbarTeamRead(d, meta)
}

func resourceRollbarTeamRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API

	team, _, getErr := client.Teams.Get(StringToInt(d.Id()))
	if getErr != nil {
		return getErr
	}

	var setErr error
	setErr = d.Set("name", team.GetResult().GetName())
	setErr = d.Set("access_level", team.GetResult().GetAccessLevel())
	setErr = d.Set("account_id", team.GetResult().GetAccountID())

	return setErr
}

func resourceRollbarTeamDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API

	log.Printf("[DEBUG] Team id to be deleted: %v", d.Id())

	_, deleteErr := client.Teams.Delete(StringToInt(d.Id()))
	if deleteErr != nil {
		return deleteErr
	}

	log.Printf("[DEBUG] Deleted team id : %v", d.Id())
	d.SetId("")

	return nil
}
