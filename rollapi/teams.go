package rollapi

// TeamsService handles communication with the teams related
// methods of the Rollbar API.
//
// Rollbar API docs: https://docs.rollbar.com/reference#teams
type TeamsService service

// Team represents a team in Rollbar.
type Team struct {
	ID          *int64  `json:"id,omitempty"`
	AccountID   *int64  `json:"account_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	AccessLevel *string `json:"access_level,omitempty"`
}

// TeamResponse represents the response returned after getting a team.
type TeamResponse struct {
	ErrorCount int   `json:"err,omitempty"`
	Result     *Team `json:"result,omitempty"`
}

// TeamListResponse represents the response returned after getting all teams.
type TeamListResponse struct {
	ErrorCount int     `json:"err,omitempty"`
	Result     []*Team `json:"result,omitempty"`
}

// TeamRequest represents a request to create a team.
type TeamRequest struct {
	Name        string `json:"name,omitempty"`
	AccessLevel string `json:"access_level,omitempty"`
}

// List all teams.
//
// Rollbar API docs: https://docs.rollbar.com/reference#list-all-teams
func (t *TeamsService) List() (*TeamListResponse, *Response, error) {
	var result *TeamListResponse
	urlStr := t.client.requestURL("/teams")

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Create a team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#create-a-team
func (t *TeamsService) Create(opts *TeamRequest) (*TeamResponse, *Response, error) {
	var result *TeamResponse
	urlStr := t.client.requestURL("/teams")

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Post(urlStr, &result, opts)

	return result, response, getErr
}

// Get a single teams.
//
// Rollbar API docs: https://docs.rollbar.com/reference#get-a-team
func (t *TeamsService) Get(teamID int) (*TeamResponse, *Response, error) {
	var result *TeamResponse
	urlStr := t.client.requestURL("/team/%d", teamID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Delete an existing project.
//
// Rollbar API docs: https://docs.rollbar.com/reference#delete-a-team
func (t *TeamsService) Delete(teamID int) (*Response, error) {
	urlStr := t.client.requestURL("/team/%d", teamID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Delete(urlStr, nil)

	return response, getErr
}
