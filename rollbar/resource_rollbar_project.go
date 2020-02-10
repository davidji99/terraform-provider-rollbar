package rollbar

import (
	"github.com/davidji99/terraform-provider-rollbar/rollapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"regexp"
)

func resourceRollbarProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceRollbarProjectCreate,
		Read:   resourceRollbarProjectRead,
		Delete: resourceRollbarProjectDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRollbarProjectImport,
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

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceRollbarProjectImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	readErr := resourceRollbarProjectRead(d, meta)

	return []*schema.ResourceData{d}, readErr
}

func resourceRollbarProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API
	opts := &rollapi.ProjectRequest{}

	if v, ok := d.GetOk("name"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] project name is : %s", vs)
		opts.Name = vs
	}

	log.Printf("Creating new project %s", opts.Name)

	newProject, _, createErr := client.Projects.Create(opts)
	if createErr != nil {
		return createErr
	}

	log.Printf("Created new project %s", opts.Name)

	d.SetId(Int64ToString(newProject.GetResult().GetID()))

	return resourceRollbarProjectRead(d, meta)
}

func resourceRollbarProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API

	project, _, getErr := client.Projects.Get(StringToInt(d.Id()))
	if getErr != nil {
		return getErr
	}

	var setErr error
	setErr = d.Set("name", project.GetResult().GetName())
	setErr = d.Set("status", project.GetResult().GetStatus())
	setErr = d.Set("account_id", project.GetResult().GetAccountID())

	return setErr
}

func resourceRollbarProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API

	log.Printf("Project id to be deleted: %v", d.Id())

	_, deleteErr := client.Projects.Delete(StringToInt(d.Id()))
	if deleteErr != nil {
		return deleteErr
	}

	log.Printf("[DEBUG] Deleted Project id : %v", d.Id())
	d.SetId("")

	return nil
}
