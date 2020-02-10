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
	ErrorCount *int  `json:"err,omitempty"`
	Result     *Team `json:"result,omitempty"`
}

// TeamListResponse represents the response returned after getting all teams.
type TeamListResponse struct {
	ErrorCount *int    `json:"err,omitempty"`
	Result     []*Team `json:"result,omitempty"`
}

// TeamRequest represents a request to create a team.
type TeamRequest struct {
	Name        string `json:"name,omitempty"`
	AccessLevel string `json:"access_level,omitempty"`
}

// TeamProjectAssoc represents a team and project relationship.
type TeamProjectAssoc struct {
	TeamID    *int64 `json:"team_id,omitempty"`
	ProjectID *int64 `json:"project_id,omitempty"`
}

// TeamProjectAssocListResponse represents a response when getting all of a team's projects.
type TeamProjectAssocListResponse struct {
	ErrorCount *int                `json:"err,omitempty"`
	Result     []*TeamProjectAssoc `json:"result,omitempty"`
}

// TeamProjectAssocListResponse represents a response when getting  a team's project.
type TeamProjectAssocResponse struct {
	ErrorCount *int              `json:"err,omitempty"`
	Result     *TeamProjectAssoc `json:"result,omitempty"`
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

// ListUsers all users for a team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#list-a-teams-users
func (t *TeamsService) ListUsers(teamID int) (*UserListResponse, *Response, error) {
	var result *UserListResponse
	urlStr := t.client.requestURL("/team/%d/users", teamID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// IsUserAMember checks if a user is assigned to a team. Returns true if user is a member, false otherwise.
//
// Rollbar API docs: https://docs.rollbar.com/reference#check-if-a-user-is-assigned-to-a-team
func (t *TeamsService) IsUserAMember(teamID, userID int) (bool, *Response, error) {
	isMember := false
	urlStr := t.client.requestURL("/team/%d/user/%d", teamID, userID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Get(urlStr, nil, nil)
	if getErr != nil {
		return false, nil, getErr
	}

	// Per API documentation, the response returns a 200 if user belongs to the team
	if response.StatusCode == 200 {
		isMember = true
	}

	return isMember, response, nil
}

// AddUser assigns a user to team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#assign-a-user-to-team
func (t *TeamsService) AddUser(teamID, userID int) (bool, *Response, error) {
	urlStr := t.client.requestURL("/team/%d/user/%d", teamID, userID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Put(urlStr, nil, nil)
	if getErr != nil {
		return false, nil, getErr
	}

	return true, response, nil
}

// RemoveUser removes a user to team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#remove-a-user-from-a-team
func (t *TeamsService) RemoveUser(teamID, userID int) (bool, *Response, error) {
	urlStr := t.client.requestURL("/team/%d/user/%d", teamID, userID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Delete(urlStr, nil)
	if getErr != nil {
		return false, nil, getErr
	}

	return true, response, nil
}

// TeamInviteRequest represents a request to invite a user to a team.
type TeamInviteRequest struct {
	Email string `json:"email,omitempty"`
}

// InviteUser invites a user to the specific team, using the user's email address.
//
// If the email address belongs to an existing Rollbar user, they will be immediately added to the team,
// and sent an email notification. Otherwise, an invite email will be sent,
// containing a signup link that will allow the recipient to join the specified team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#invite-an-email-address-to-a-team
func (t *TeamsService) InviteUser(teamID int, opts *TeamInviteRequest) (*InvitationResponse, *Response, error) {
	var result *InvitationResponse
	urlStr := t.client.requestURL("/team/%d/invites", teamID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Post(urlStr, &result, opts)

	return result, response, getErr
}

// ListInvitations returns all invitations of a given team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#list-invitations-to-a-team
func (t *TeamsService) ListInvitations(teamID int) (*InvitationListResponse, *Response, error) {
	var result *InvitationListResponse
	urlStr := t.client.requestURL("/team/%d/invites", teamID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// ListProjects returns all of a team's projects.
//
// Rollbar API docs: https://docs.rollbar.com/reference#list-a-teams-projects
func (t *TeamsService) ListProjects(teamID int) (*TeamProjectAssocListResponse, *Response, error) {
	var result *TeamProjectAssocListResponse
	urlStr := t.client.requestURL("/team/%d/projects", teamID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// AssignProject assigns a project to a team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#assign-a-team-to-a-project
func (t *TeamsService) AssignProject(teamID, projectID int) (*TeamProjectAssocResponse, *Response, error) {
	var result *TeamProjectAssocResponse
	urlStr := t.client.requestURL("/team/%d/project/%d", teamID, projectID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Put(urlStr, &result, nil)

	return result, response, getErr
}

// RemoveProject remove a project from a team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#remove-a-team-from-a-project
func (t *TeamsService) RemoveProject(teamID, projectID int) (*Response, error) {
	urlStr := t.client.requestURL("/team/%d/project/%d", teamID, projectID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Delete(urlStr, nil)

	return response, getErr
}

// HasProject checks if a project is assigned to a team.
//
// Rollbar API docs: https://docs.rollbar.com/reference#check-if-a-team-is-assigned-to-a-project
func (t *TeamsService) HasProject(teamID, projectID int) (bool, *Response, error) {
	var result *TeamProjectAssocResponse
	urlStr := t.client.requestURL("/team/%d/project/%d", teamID, projectID)

	// Set the correct authentication header
	t.client.setAuthTokenHeader(t.client.accountAccessToken)

	// Execute the request
	response, getErr := t.client.Get(urlStr, &result, nil)
	if getErr != nil {
		return false, response, getErr
	}

	return true, response, nil
}
