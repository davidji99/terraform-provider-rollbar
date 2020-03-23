package rollbar

import (
	"encoding/json"
	"fmt"
	"github.com/davidji99/rollrest-go/rollrest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"strconv"
	"time"
)

var (
	validTriggers = []string{"new_item", "occurrence_rate", "resolved_item",
		"reactivated_item", "exp_repeat_item"}

	validFilters = []string{"environment", "level", "title", "filename",
		"context", "method", "framework", "path", "rate",
		"unique_occurrences"}
)

func resourceRollbarPagerDutyNotificationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRollbarPagerDutyNotificationRuleCreate,
		Read:   resourceRollbarPagerDutyNotificationRuleRead,
		Update: resourceRollbarPagerDutyNotificationRuleUpdate,
		Delete: resourceRollbarPagerDutyNotificationRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRollbarPagerDutyNotificationRuleImport,
		},

		Schema: map[string]*schema.Schema{
			"rule": {
				ConfigMode: schema.SchemaConfigModeBlock,
				Type:       schema.TypeList,
				Required:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trigger": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(validTriggers, false),
						},

						"filter": {
							ConfigMode: schema.SchemaConfigModeBlock,
							Type:       schema.TypeList,
							Required:   true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(validFilters, false),
									},

									"operation": {
										Type:     schema.TypeString,
										Required: true,
									},

									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"period": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"count": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"config": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_key": {
										Type:         schema.TypeString,
										Required:     true, // as it is the only one defined.
										ValidateFunc: validation.StringLenBetween(32, 32),
									},
								},
							},
						},
					},
				},
			}, // end of rule block
		},
	}
}

func resourceRollbarPagerDutyNotificationRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return nil, fmt.Errorf("not possible to import PagerDuty notification rule(s) due to API limitations")
}

func resourceRollbarPagerDutyNotificationRuleCreate(d *schema.ResourceData, meta interface{}) error {
	// The Create function will call the Update function since the API does not have a Post, only Put.
	updateErr := resourceRollbarPagerDutyNotificationRuleUpdate(d, meta)
	if updateErr != nil {
		return updateErr
	}

	// Set the resource ID to be the epoch time in nanoseconds
	d.SetId(strconv.Itoa(time.Now().Nanosecond()))

	return resourceRollbarPagerDutyNotificationRuleUpdate(d, meta)
}

func resourceRollbarPagerDutyNotificationRuleRead(d *schema.ResourceData, meta interface{}) error {
	// The API doesn't support any GET endpoints for this resource so just setting what is defined into state
	out, marshallErr := json.Marshal(constructRuleDefinitions(d))
	if marshallErr != nil {
		return marshallErr
	}

	ruleListMap := make([]map[string]interface{}, 0)
	err := json.Unmarshal(out, &ruleListMap)
	if err != nil {
		return err
	}

	// Rename all map keys named 'filters' in ruleListMap to 'filter' to be consistent with resource schema before
	// saving to state. The reason being the JSON schema for this resource's request requires 'filters'.
	for _, ruleMap := range ruleListMap {
		for k, v := range ruleMap {
			if k == "filters" {
				ruleMap["filter"] = v
				delete(ruleMap, k)
			}
		}
	}

	log.Printf("[DEBUG] Rules to be stored in state %v", ruleListMap)

	return d.Set("rule", ruleListMap)
}

func resourceRollbarPagerDutyNotificationRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API
	opts := constructRuleDefinitions(d)

	log.Printf("[DEBUG] Modifying PagerDuty notification rule(s) %v", opts)

	isModified, _, modifyErr := client.Notifications.ModifyPagerDutyRules(opts)
	if modifyErr != nil {
		return modifyErr
	}

	log.Printf("[DEBUG] Was modifying PagerDuty notification rule(s) successful: %v", isModified)

	return resourceRollbarPagerDutyNotificationRuleRead(d, meta)
}

func resourceRollbarPagerDutyNotificationRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).API

	log.Printf("[DEBUG] Deleting all PagerDuty notification rules")

	isDeleted, _, deleteErr := client.Notifications.DeleteAllPagerDutyRules()
	if deleteErr != nil {
		return deleteErr
	}

	log.Printf("[DEBUG] Was all PagerDuty notification rules deleted: %v", isDeleted)

	d.SetId("")

	return nil
}

// constructRuleDefinitions returns a properly formatted and nested notification rule request.
func constructRuleDefinitions(d *schema.ResourceData) []*rollrest.PDRuleRequest {
	opts := make([]*rollrest.PDRuleRequest, 0)

	if ruleListRaw, ok := d.GetOk("rule"); ok {
		ruleList := ruleListRaw.([]interface{})

		// Iterate through the ruleList
		for _, ruleRaw := range ruleList {
			pdRule := &rollrest.PDRuleRequest{}

			rule := ruleRaw.(map[string]interface{})

			// Define trigger
			if triggerRaw, ok := rule["trigger"]; ok {
				pdRule.Trigger = triggerRaw.(string)
			}

			// Define config
			if configRaw, ok := rule["config"]; ok {
				config := configRaw.(map[string]interface{})
				configOpt := &rollrest.PDRuleConfig{}

				if serviceKeyRaw, ok := config["service_key"]; ok {
					configOpt.ServiceKey = serviceKeyRaw.(string)
				}

				pdRule.Config = configOpt
			}

			// Define filters
			if ruleFilterListRaw, ok := rule["filter"]; ok {
				ruleFilterOpts := make([]*rollrest.PDRuleFilter, 0)
				filterList := ruleFilterListRaw.([]interface{})

				// Iterate through the filters
				for _, ruleFilterRaw := range filterList {
					ruleFilterOpt := &rollrest.PDRuleFilter{}
					ruleFilter := ruleFilterRaw.(map[string]interface{})

					if typeRaw, ok := ruleFilter["type"]; ok {
						ruleFilterOpt.Type = typeRaw.(string)
					}

					if operationRaw, ok := ruleFilter["operation"]; ok {
						ruleFilterOpt.Operation = operationRaw.(string)
					}

					if valueRaw, ok := ruleFilter["value"]; ok {
						ruleFilterOpt.Value = valueRaw.(string)
					}

					if pathRaw, ok := ruleFilter["path"]; ok {
						ruleFilterOpt.Path = pathRaw.(string)
					}

					if periodRaw, ok := ruleFilter["period"]; ok {
						ruleFilterOpt.Period = periodRaw.(int)
					}

					if countRaw, ok := ruleFilter["count"]; ok {
						ruleFilterOpt.Count = countRaw.(int)
					}

					// Add the new ruleFilterOpt to ruleFilterOpts
					ruleFilterOpts = append(ruleFilterOpts, ruleFilterOpt)
				}

				// Add ruleFilterOpts to the pdRule
				pdRule.Filters = ruleFilterOpts
			}

			// Finally, add the rule to opts
			opts = append(opts, pdRule)
		}
	} // end of defining rule

	return opts
}
