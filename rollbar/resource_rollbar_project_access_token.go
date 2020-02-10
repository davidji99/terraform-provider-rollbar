package rollbar

import (
	"fmt"
	"github.com/davidji99/terraform-provider-rollbar/rollapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"regexp"
)

func resourceRollbarProjectAccessToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceRollbarProjectAccessTokenCreate,
		Read:   resourceRollbarProjectAccessTokenRead,
		Update: resourceRollbarProjectAccessTokenUpdate,
		Delete: resourceRollbarProjectAccessTokenDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRollbarProjectAccessTokenImport,
		},

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z][0-9A-Za-z,.\-_]{1,31}$`),
					"Must start with a letter and can only contain letters, numbers, underscores, "+
						"hyphens, and periods. Max length 32 characters."),
			},

			"scopes": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice(
						[]string{"read", "write", "post_server_item", "post_client_server"}, false),
				},
			},

			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
			},

			"rate_limit_window_size": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(1),
			},

			"rate_limit_window_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(1),
			},

			"cur_rate_limit_window_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"date_created": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"access_token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceRollbarProjectAccessTokenImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// To import this resource, users must pass in the project ID & access token as the 'ID'.
	// We then proceed to set one half of the import ID as the "access_token"
	// before generating a random string number to set as the real resource ID in state.
	projectId, accessToken, parseErr := ParseCompositeID(d.Id())
	if parseErr != nil {
		return nil, parseErr
	}

	var setErr error
	setErr = d.Set("access_token", accessToken)
	setErr = d.Set("project_id", StringToInt(projectId))
	if setErr != nil {
		return nil, setErr
	}

	d.SetId(GenerateRandomResourceID())

	readErr := resourceRollbarProjectAccessTokenRead(d, meta)

	return []*schema.ResourceData{d}, readErr
}

func resourceRollbarProjectAccessTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API
	opts := &rollapi.PATCreateRequest{}

	if v, ok := d.GetOk("name"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project access token name : %s", vs)
		opts.Name = vs
	}

	if scopes, ok := d.GetOk("scopes"); ok {
		s := scopes.(*schema.Set).List()
		scopesToAdd := make([]string, 0)

		for _, scope := range s {
			scopesToAdd = append(scopesToAdd, scope.(string))
		}
		opts.Scopes = scopesToAdd
	}

	if v, ok := d.GetOk("status"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project access token status : %s", vs)
		opts.Status = vs
	}

	if v, ok := d.GetOk("rate_limit_window_size"); ok {
		vs := v.(int)

		// On creation for a new project access token, the API only accepts certain
		// values for rate_limit_window_size. Therefore, we will validate the user value here.
		if valErr := validateWinSizeOnCreation(vs); valErr != nil {
			return valErr
		}

		log.Printf("[DEBUG] project access token rate_limit_window_size : %d", vs)
		opts.RateLimitWindowSize = &vs
	}

	if v, ok := d.GetOk("rate_limit_window_count"); ok {
		vs := v.(int)
		log.Printf("[DEBUG] project access token RateLimitWindowCount : %d", vs)
		opts.RateLimitWindowCount = &vs
	}

	log.Printf("Creating project access token %s", opts.Name)

	newPAT, _, createErr := client.ProjectAccessTokens.Create(getProjectID(d), opts)
	if createErr != nil {
		return createErr
	}

	log.Printf("Created project access token %s", opts.Name)

	// Set the ID of the resource to a random number as there's no unique ID for this resource remotely
	// and we don't want to use the access token as a state resource ID.
	d.SetId(GenerateRandomResourceID())

	// Set the access token value so we can use the value in the READ function.
	d.Set("access_token", newPAT.GetResult().GetAccessToken())

	return resourceRollbarProjectAccessTokenRead(d, meta)
}

func validateWinSizeOnCreation(i int) error {
	supportedWinSizeOnCreate := []int{0, 60, 300, 1800, 3600, 86400, 604800, 2592000}
	for _, v := range supportedWinSizeOnCreate {
		if v == i {
			return nil
		}
	}

	return fmt.Errorf("%d is not a supported window size for token creation. "+
		"Valid values are: %+v", i, supportedWinSizeOnCreate)
}

func resourceRollbarProjectAccessTokenRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API

	projectID := getProjectID(d)

	pat, _, getErr := client.ProjectAccessTokens.Get(projectID, getAccessToken(d))
	if getErr != nil {
		return getErr
	}

	var setErr error
	setErr = d.Set("project_id", pat.GetProjectID())
	setErr = d.Set("name", pat.GetName())
	setErr = d.Set("scopes", pat.Scopes)
	setErr = d.Set("status", pat.GetStatus())
	setErr = d.Set("rate_limit_window_size", pat.GetRateLimitWindowSize())
	setErr = d.Set("rate_limit_window_count", pat.GetRateLimitWindowCount())
	setErr = d.Set("cur_rate_limit_window_count", pat.GetCurrentRateLimitWindowCount())
	setErr = d.Set("date_created", int(pat.GetDataCreated()))
	setErr = d.Set("access_token", pat.GetAccessToken())

	return setErr
}

func resourceRollbarProjectAccessTokenUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API
	opts := &rollapi.PATUpdateRequest{}

	if v, ok := d.GetOk("rate_limit_window_size"); ok {
		vs := v.(int)
		log.Printf("[DEBUG] project access token rate_limit_window_size : %d", vs)
		opts.RateLimitWindowSize = &vs
	}

	if v, ok := d.GetOk("rate_limit_window_count"); ok {
		vs := v.(int)
		log.Printf("[DEBUG] project access token rate_limit_window_count : %d", vs)
		opts.RateLimitWindowCount = &vs
	}
	pat, _, updateErr := client.ProjectAccessTokens.Update(getProjectID(d), getAccessToken(d), opts)
	if updateErr != nil {
		return updateErr
	}

	log.Printf("Updated project access token %s", pat.GetResult().GetName())

	return resourceRollbarProjectAccessTokenRead(d, meta)
}

func resourceRollbarProjectAccessTokenDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] There is no API DELETE support for project access token resource so this is a no-op. " +
		"Resource will be removed only from state.")
	d.SetId("")

	return nil
}

// getProjectID is a helper method to get the project id.
func getProjectID(d *schema.ResourceData) int {
	var projectID int
	if v, ok := d.GetOk("project_id"); ok {
		vs := v.(int)
		log.Printf("[DEBUG] Project id: %d", vs)
		projectID = vs
	}

	return projectID
}

// getAccessToken is a helper method to get the access token.
//
// This should only be used for this resource's update function.
func getAccessToken(d *schema.ResourceData) string {
	var accessToken string
	if v, ok := d.GetOk("access_token"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] access_token: %s", vs)
		accessToken = vs
	}

	return accessToken
}
