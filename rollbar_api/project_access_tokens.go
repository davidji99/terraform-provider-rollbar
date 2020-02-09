package rollbar_api

import "fmt"

// ProjectAccessTokensService handles communication with the project access token related
// methods of the Rollbar API.
//
// Rollbar API docs: https://docs.rollbar.com/reference#project-access-tokens
type ProjectAccessTokensService service

// ProjectAccessToken represents a project access token.
type ProjectAccessToken struct {
	ProjectID                   *int64   `json:"project_id,omitempty"`
	AccessToken                 *string  `json:"access_token,omitempty"`
	Name                        *string  `json:"name,omitempty"`
	Status                      *string  `json:"status,omitempty"`
	RateLimitWindowSize         *int64   `json:"rate_limit_window_size,omitempty"`
	RateLimitWindowCount        *int64   `json:"rate_limit_window_count,omitempty"`
	CurrentRateLimitWindowStart *int64   `json:"cur_rate_limit_window_start,omitempty"`
	CurrentRateLimitWindowCount *int64   `json:"cur_rate_limit_window_count,omitempty"`
	DataCreated                 *int64   `json:"date_created,omitempty"`
	DateModified                *int64   `json:"date_modified,omitempty"`
	Scopes                      []string `json:"scopes,omitempty"`
}

type ProjectAccessTokenResponse struct {
	ErrorCount int                 `json:"err,omitempty"`
	Result     *ProjectAccessToken `json:"result,omitempty"`
}

// ProjectAccessTokenListResponse represents the response returned after getting all project access tokens.
type ProjectAccessTokenListResponse struct {
	ErrorCount int                   `json:"err,omitempty"`
	Result     []*ProjectAccessToken `json:"result,omitempty"`
}

// PATCreateRequest represents a request to create a project access token.
type PATCreateRequest struct {
	Name                 string   `json:"name,omitempty"`
	Scopes               []string `json:"scopes,omitempty"`
	Status               string   `json:"status,omitempty"`
	RateLimitWindowSize  *int     `json:"rate_limit_window_size,omitempty"`
	RateLimitWindowCount *int     `json:"rate_limit_window_count,omitempty"`
}

// PATUpdateRequest represents a request to update a project access token.
type PATUpdateRequest struct {
	RateLimitWindowSize  *int `json:"rate_limit_window_size,omitempty"`
	RateLimitWindowCount *int `json:"rate_limit_window_count,omitempty"`
}

// List all of a project's access tokens.
//
// Rollbar API docs: https://docs.rollbar.com/reference#list-all-project-access-tokens
func (p *ProjectAccessTokensService) List(projectID int) (*ProjectAccessTokenListResponse, *Response, error) {
	var result *ProjectAccessTokenListResponse
	urlStr := p.client.requestURL("/project/%d/access_tokens", projectID)

	// Set the correct authentication header
	p.client.setAuthTokenHeader(p.client.accountAccessToken)

	// Execute the request
	response, getErr := p.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Get a single project access tokens using the date_created value.
//
// We don't want to use the actual access token.
//
// Also as no endpoint officially exists, this function will first fetch all of a project's access token
// and iterate through each token to find the specified one.
func (p *ProjectAccessTokensService) Get(projectID int, accessToken string) (*ProjectAccessToken, *Response, error) {
	projects, _, listErr := p.List(projectID)
	if listErr != nil {
		return nil, nil, listErr
	}

	var targetProject *ProjectAccessToken

	for _, project := range projects.Result {
		if project.GetAccessToken() == accessToken {
			targetProject = project
		}
	}

	if targetProject == nil {
		return nil, nil, fmt.Errorf("not found")
	}

	return targetProject, nil, nil
}

// Create a project access token.
//
// Rollbar API docs: https://docs.rollbar.com/reference#create-a-project-access-token
func (p *ProjectAccessTokensService) Create(projectID int, opts *PATCreateRequest) (*ProjectAccessTokenResponse, *Response, error) {
	var result *ProjectAccessTokenResponse
	urlStr := p.client.requestURL("/project/%d/access_tokens", projectID)

	// Set the correct authentication header
	p.client.setAuthTokenHeader(p.client.accountAccessToken)

	// Execute the request
	response, getErr := p.client.Post(urlStr, &result, opts)

	return result, response, getErr
}

// Update a project access token.
//
// Rollbar API docs: https://docs.rollbar.com/reference#update-a-rate-limit
func (p *ProjectAccessTokensService) Update(projectID int, accessToken string,
	opts *PATUpdateRequest) (*ProjectAccessTokenResponse, *Response, error) {
	// API requires RateLimitWindowSize and RateLimitWindowCount to be both set in the request body so validate this first.
	if opts.RateLimitWindowSize == nil || opts.RateLimitWindowCount == nil {
		return nil, nil, fmt.Errorf("both rate_limit_window_size & rate_limit_window_count " +
			"must be set in the request body")
	}

	var result *ProjectAccessTokenResponse
	urlStr := p.client.requestURL("/project/%d/access_token/%s", projectID, accessToken)

	// Set the correct authentication header
	p.client.setAuthTokenHeader(p.client.accountAccessToken)

	// Execute the request
	response, getErr := p.client.Patch(urlStr, &result, opts)

	return result, response, getErr
}

// TODO: add support for deleting project access tokens when the DELETE endpoint is available.
