package rollbar

import (
	"encoding/json"
	"fmt"
	"github.com/davidji99/terraform-provider-rollbar/rollapi"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"regexp"
	"strconv"
	"time"
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
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringMatch(
								regexp.MustCompile(`^new_item$`), "only valid value is 'new_item'"),
						},

						"filter": {
							ConfigMode: schema.SchemaConfigModeBlock,
							Type:       schema.TypeList,
							Required:   true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice(
											[]string{"environment", "level", "title", "filename",
												"context", "method", "framework", "path"}, false),
									},

									"operation": {
										Type:     schema.TypeString,
										Required: true,
									},

									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"path": {
										Type:     schema.TypeString,
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
	// saving to state.
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
func constructRuleDefinitions(d *schema.ResourceData) []*rollapi.PDRuleRequest {
	opts := make([]*rollapi.PDRuleRequest, 0)

	if ruleListRaw, ok := d.GetOk("rule"); ok {
		ruleList := ruleListRaw.([]interface{})

		// Iterate through the ruleList
		for _, ruleRaw := range ruleList {
			pdRule := &rollapi.PDRuleRequest{}

			rule := ruleRaw.(map[string]interface{})

			// Define trigger
			if triggerRaw, ok := rule["trigger"]; ok {
				pdRule.Trigger = triggerRaw.(string)
			}

			// Define config
			if configRaw, ok := rule["config"]; ok {
				config := configRaw.(map[string]interface{})
				configOpt := &rollapi.PDRuleConfig{}

				if serviceKeyRaw, ok := config["service_key"]; ok {
					configOpt.ServiceKey = serviceKeyRaw.(string)
				}

				pdRule.Config = configOpt
			}

			// Define filters
			if ruleFilterListRaw, ok := rule["filter"]; ok {
				ruleFilterOpts := make([]*rollapi.PDRuleFilter, 0)
				filterList := ruleFilterListRaw.([]interface{})

				// Iterate through the filters
				for _, ruleFilterRaw := range filterList {
					ruleFilterOpt := &rollapi.PDRuleFilter{}
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
