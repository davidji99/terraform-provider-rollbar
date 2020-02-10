package rollapi

// NotificationsService handles communication with the notification related
// methods of the Rollbar API.
//
// Rollbar API docs: https://docs.rollbar.com/reference#notifications
type NotificationsService service

// PDIntegrationRequest represents a request to configure Rollbar with PagerDuty.
type PDIntegrationRequest struct {
	Enabled    *bool  `json:"enabled,omitempty"`
	ServiceKey string `json:"service_key,omitempty"`
}

// PDRuleRequest represents a request to add one or many PagerDuty notification rule.
type PDRuleRequest struct {
	// As of Feb. 11th 2020, the only possible value for the Triggers field is `new_item`.
	Trigger string          `json:"trigger,omitempty"`
	Filters []*PDRuleFilter `json:"filters,omitempty"`
	Config  *PDRuleConfig   `json:"config,omitempty"`
}

// PDRuleFilter represents a PagerDuty rule filter.
type PDRuleFilter struct {
	Type      string `json:"type,omitempty"`
	Operation string `json:"operation,omitempty"`
	Value     string `json:"value,omitempty"`
}

// PDRuleConfig represents the configuration options available on a rule.
type PDRuleConfig struct {
	// PagerDuty Service API Key. Make sure the ServiceKey string value is length 32.
	ServiceKey string `json:"service_key,omitempty"`
}

// ConfigurePagerDutyIntegration configures the PagerDuty integration for a project.
//
// This function creates and modifies the PagerDuty integration. Requires a project access token.
//
// Rollbar API docs: https://docs.rollbar.com/reference#configuring-pagerduty-integration
func (n *NotificationsService) ConfigurePagerDutyIntegration(opts *PDIntegrationRequest) (*Response, error) {
	urlStr := n.client.requestURL("/notifications/pagerduty")

	// Set the correct authentication header
	n.client.setAuthTokenHeader(n.client.projectAccessToken)

	// Execute the request
	response, getErr := n.client.Put(urlStr, nil, opts)

	return response, getErr
}

// ModifyPagerDutyRules creates & modifies PagerDuty notification rules for a project.
//
// Requires a project access token.
// (The API documentation is wrong regarding which documentation to use as of Feb. 10th, 2020.)
//
// Rollbar API docs: https://docs.rollbar.com/reference#setup-pagerduty-notification-rules
func (n *NotificationsService) ModifyPagerDutyRules(opts []*PDRuleRequest) (bool, *Response, error) {
	urlStr := n.client.requestURL("/notifications/pagerduty/rules")

	// Set the correct authentication header
	n.client.setAuthTokenHeader(n.client.projectAccessToken)

	// Execute the request
	isSuccessful := false
	response, getErr := n.client.Put(urlStr, nil, opts)
	if getErr != nil {
		return isSuccessful, response, getErr
	}

	if response.StatusCode == 200 {
		isSuccessful = true
	}

	return isSuccessful, response, nil
}

// DeleteAllPagerDutyRules removes all rules for a project's PagerDuty notification integration.
// This is the same ModifyPagerDutyRules but passes in an empty array as the request body for convenience.
//
// Requires a project access token.
// (The API documentation is wrong regarding which documentation to use as of Feb. 10th, 2020.)
//
// Rollbar API docs: https://docs.rollbar.com/reference#setup-pagerduty-notification-rules
func (n *NotificationsService) DeleteAllPagerDutyRules() (bool, *Response, error) {
	opts := make([]*PDRuleRequest, 0)
	return n.ModifyPagerDutyRules(opts)
}
