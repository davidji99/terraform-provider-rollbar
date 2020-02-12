package rollapi

// ProjectsService handles communication with the project related
// methods of the Rollbar API.
//
// Rollbar API docs: https://docs.rollbar.com/reference#projects
type ProjectsService service

// Project represents a rollbar project.
type Project struct {
	ID           *int64     `json:"id,omitempty"`
	AccountID    *int64     `json:"account_id,omitempty"`
	Status       *string    `json:"status,omitempty"`
	DataCreated  *int64     `json:"date_created,omitempty"`
	DateModified *int64     `json:"date_modified,omitempty"`
	Name         *string    `json:"name,omitempty"`
	SettingsData *ProjectSD `json:"settings_data,omitempty"`
}

// ProjectResponse represents the response returned after a successful GET/POST.
type ProjectResponse struct {
	ErrorCount *int     `json:"err,omitempty"`
	Result     *Project `json:"result,omitempty"`
}

// ProjectListResponse represents the response returned after getting all projects.
type ProjectListResponse struct {
	ErrorCount *int       `json:"err,omitempty"`
	Results    []*Project `json:"result,omitempty"`
}

// ProjectSD represents a project's settings data.
type ProjectSD struct {
	FingerprintVersions *ProjectSDFingerprintVersions `json:"fingerprint_versions,omitempty"`
	Migrations          *ProjectSDMigrations          `json:"migrations,omitempty"`
}

// ProjectSDFingerprintVersions represents a project settings data's fingerprint versions.
type ProjectSDFingerprintVersions struct {
	BrowserBrowserJS    *int `json:"browser.browser-js,omitempty"`
	AndroidAndroid      *int `json:"android.android,omitempty"`
	UnminifyReactErrors *int `json:"unminify_react_errors,omitempty"`
}

// ProjectSDMigrations represents a project settings data's migrations.
type ProjectSDMigrations struct {
	EnableSourceMaps                   *int `json:"enable_source_maps,omitempty"`
	EnableCustomFingerprintingOverride *int `json:"enable_custom_fingerprinting_override,omitempty"`
	RecognizeDirectRecursion           *int `json:"recognize_direct_recursion,omitempty"`
	EnableMissingJquery                *int `json:"enable_missing_jquery,omitempty"`
	EnableCalculateSymbolRanges        *int `json:"enable_calculate_symbol_ranges,omitempty"`
}

// ProjectRequest represents a request to create a project.
//
// Currently, it is not possible to update an existing project via the API.
type ProjectRequest struct {
	Name string `json:"name,omitempty"`
}

// List all non-deleted projects.
//
// By default, the API returns all a list of deleted and active projects. If you wish to see deleted projects,
// please use the ListAll() function.
//
// Rollbar API docs: https://docs.rollbar.com/reference#list-all-projects
func (p *ProjectsService) List() (*ProjectListResponse, *Response, error) {
	var result *ProjectListResponse
	urlStr := p.client.requestURL("/projects")

	// Set the correct authentication header
	p.client.setAuthTokenHeader(p.client.accountAccessToken)

	// Execute the request
	response, getErr := p.client.Get(urlStr, &result, nil)
	if getErr != nil {
		return nil, nil, getErr
	}

	// If there are any results, iterate through them and get only the active projects.
	if len(result.Results) > 0 {
		var activeProjects []*Project
		for _, project := range result.Results {
			if project.GetName() != "" {
				activeProjects = append(activeProjects, project)
			}
		}

		result.Results = activeProjects
	}

	return result, response, nil
}

// List all projects, including deleted ones.
//
// By default, the API returns all a list of deleted and active projects. If you wish to see only active projects,
// please use the List() function.
//
// Rollbar API docs: https://docs.rollbar.com/reference#list-all-projects
func (p *ProjectsService) ListAll() (*ProjectListResponse, *Response, error) {
	var result *ProjectListResponse
	urlStr := p.client.requestURL("/projects")

	// Set the correct authentication header
	p.client.setAuthTokenHeader(p.client.accountAccessToken)

	// Execute the request
	response, getErr := p.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Get a single project.
//
// Rollbar API docs: https://docs.rollbar.com/reference#get-a-project
func (p *ProjectsService) Get(id int) (*ProjectResponse, *Response, error) {
	var result *ProjectResponse
	urlStr := p.client.requestURL("/project/%d", id)

	// Set the correct authentication header
	p.client.setAuthTokenHeader(p.client.accountAccessToken)

	// Execute the request
	response, getErr := p.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Create a single project.
//
// Rollbar API docs: https://docs.rollbar.com/reference#create-a-project
func (p *ProjectsService) Create(opts *ProjectRequest) (*ProjectResponse, *Response, error) {
	var result *ProjectResponse
	urlStr := p.client.requestURL("/projects")

	// Set the correct authentication header
	p.client.setAuthTokenHeader(p.client.accountAccessToken)

	// Execute the request
	response, getErr := p.client.Post(urlStr, &result, opts)

	return result, response, getErr
}

// Delete an existing project.
//
// Rollbar API docs: https://docs.rollbar.com/reference#delete-a-project
func (p *ProjectsService) Delete(id int) (*Response, error) {
	urlStr := p.client.requestURL("/project/%d", id)

	// Set the correct authentication header
	p.client.setAuthTokenHeader(p.client.accountAccessToken)

	// Execute the request
	response, getErr := p.client.Delete(urlStr, nil)

	return response, getErr
}
