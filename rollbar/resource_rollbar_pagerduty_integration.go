package rollbar

import (
	"fmt"
	"github.com/davidji99/terraform-provider-rollbar/rollapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"strconv"
	"time"
)

func resourceRollbarPagerDutyIntegration() *schema.Resource {
	return &schema.Resource{
		Create: resourceRollbarPagerDutyIntegrationCreate,
		Read:   resourceRollbarPagerDutyIntegrationRead,
		Delete: resourceRollbarPagerDutyIntegrationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRollbarPagerDutyIntegrationImport,
		},

		Schema: map[string]*schema.Schema{
			"service_key": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(32, 32),
			},

			"enabled": {
				Type:     schema.TypeBool,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceRollbarPagerDutyIntegrationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return nil, fmt.Errorf("not possible to import PagerDuty integration due to API limitations")
}

func resourceRollbarPagerDutyIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API
	opts := &rollapi.PDIntegrationRequest{}

	if v, ok := d.GetOk("service_key"); ok {
		vs := v.(string)
		log.Printf("[DEBUG] service_key is : %s", vs)
		opts.ServiceKey = vs
	}

	vs := d.Get("enabled").(bool)
	log.Printf("[DEBUG] enabled is : %v", vs)
	opts.Enabled = &vs

	log.Printf("[DEBUG] Adding pagerduty integration %v", opts)

	_, createErr := client.Notifications.ConfigurePagerDutyIntegration(opts)
	if createErr != nil {
		return createErr
	}

	log.Printf("[DEBUG] Added pagerduty integration %v", opts)

	//if d.Get("delete_default_rules").(bool) {
	//	log.Printf("[DEBUG] Deleting default rules added on service_key %s", opts.ServiceKey)
	//
	//	_, _, deleteErr := client.Notifications.DeleteAllPagerDutyRules()
	//	if deleteErr != nil {
	//		return deleteErr
	//	}
	//
	//	log.Printf("[DEBUG] Deleted default rules added on service_key %s", opts.ServiceKey)
	//}

	// Set the resource ID to be the epoch time in nanoseconds
	d.SetId(strconv.Itoa(time.Now().Nanosecond()))

	return resourceRollbarPagerDutyIntegrationRead(d, meta)
}

func resourceRollbarPagerDutyIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	// There is no GET API endpoint so state will set whatever value is defined in the user's configuration.
	var setErr error
	setErr = d.Set("service_key", d.Get("service_key"))
	setErr = d.Set("enabled", d.Get("enabled"))

	return setErr
}

func resourceRollbarPagerDutyIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	// There is no DELETE API endpoint so resource deletion will entail disabling the integration.
	// Users will need to visit the UI to manually remove the integration.
	client := meta.(*Config).API
	opts := &rollapi.PDIntegrationRequest{}

	opts.ServiceKey = d.Get("service_key").(string)
	e := false
	opts.Enabled = &e

	log.Printf("[DEBUG] Disabling pagerduty integration %v", opts)

	_, createErr := client.Notifications.ConfigurePagerDutyIntegration(opts)
	if createErr != nil {
		return createErr
	}

	log.Printf("[DEBUG] Disabled pagerduty integration %v", opts)

	d.SetId("")

	return nil
}
