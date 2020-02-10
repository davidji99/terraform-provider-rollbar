package rollapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"net/url"
	"reflect"
	"sync"
	"time"
)

const (
	DefaultAPIBaseUrl = "https://api.rollbar.com/api/1"
	DefaultUserAgent  = "go-rollbar-api"
	RollbarAuthHeader = "x-rollbar-access-token"
)

// A Client manages communication with the Rollbar API.
type Client struct {
	// clientMu protects the client during calls that modify the CheckRedirect func.
	clientMu sync.Mutex

	// HTTP client used to communicate with the API.
	http *resty.Client

	// BaseURL for API
	BaseURL string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// User agent used when communicating with the Rollbar API.
	UserAgent string

	// Services used for talking to different parts of the Rollbar API.
	Projects            *ProjectsService
	ProjectAccessTokens *ProjectAccessTokensService
	Teams               *TeamsService
	Users               *UsersService

	// Custom HTTPHeaders
	customHttpHeaders map[string]string

	// Account access token
	accountAccessToken string

	// Project access token
	projectAccessToken string
}

// service represents the api service client.
type service struct {
	client *Client
}

// TokenAuthConfig represents options when initializing a new API http.
type TokenAuthConfig struct {
	// ProjectAccessToken is a Rollbar project access token.
	ProjectAccessToken *string

	// AccountAccessToken is a Rollbar account access token.
	AccountAccessToken *string

	// Custom HTTPHeaders
	CustomHttpHeaders map[string]string
}

func NewClientTokenAuth(config *TokenAuthConfig) (*Client, error) {
	// Validate that either ProjectAccessToken or AccountAccessToken are set in TokenAuthConfig.
	if config.GetProjectAccessToken() != "" && config.GetAccountAccessToken() != "" {
		return nil, fmt.Errorf("please set an account access token and/or a project access token for authentication")
	}

	// Construct new client.
	c := &Client{
		http: resty.New(), BaseURL: DefaultAPIBaseUrl, UserAgent: DefaultUserAgent,
		customHttpHeaders: config.CustomHttpHeaders, accountAccessToken: config.GetAccountAccessToken(),
		projectAccessToken: config.GetProjectAccessToken(),
	}
	c.injectServices()

	// Setup client
	c.setupClient()

	return c, nil
}

// injectServices adds the services to the client.
func (c *Client) injectServices() {
	c.common.client = c
	c.Projects = (*ProjectsService)(&c.common)
	c.ProjectAccessTokens = (*ProjectAccessTokensService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	c.Users = (*UsersService)(&c.common)
}

// setupClient sets common headers and other configurations.
func (c *Client) setupClient() {
	// We aren't setting an authentication header initially here as certain API resources require specific access_tokens.
	// Per Rollbar API documentation, each individual resource will set the access_token parameter when constructing
	// the full API endpoint URL.
	c.http.SetHeader("Content-type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", c.UserAgent).
		SetTimeout(5 * time.Minute).
		SetAllowGetMethodPayload(true)

	// Set additional headers
	if c.customHttpHeaders != nil {
		c.http.SetHeaders(c.customHttpHeaders)
	}
}

func (c *Client) setAuthTokenHeader(token string) {
	c.http.SetHeader(RollbarAuthHeader, token)
}

func (c *Client) requestURL(template string, args ...interface{}) string {
	if len(args) == 1 && args[0] == "" {
		return c.BaseURL + template
	}
	return c.BaseURL + fmt.Sprintf(template, args...)
}

// Get executes a GET http request.
func (c *Client) Get(url string, v, body interface{}) (*Response, error) {
	resp, err := c.http.R().SetResult(v).
		SetBody(body).
		Get(url)

	if err != nil {
		return nil, err
	}

	return checkResponse(resp)
}

// Post executes a POST http request.
func (c *Client) Post(url string, v, body interface{}) (*Response, error) {
	resp, err := c.http.R().SetResult(v).
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}

	return checkResponse(resp)
}

// Delete executes a DELETE http request.
func (c *Client) Delete(url string, v interface{}) (*Response, error) {
	resp, err := c.http.R().Delete(url)

	if err != nil {
		return nil, err
	}

	return checkResponse(resp)
}

// Patch executes a PATCH http request.
func (c *Client) Patch(url string, v, body interface{}) (*Response, error) {
	resp, err := c.http.R().SetResult(v).
		SetBody(body).
		Patch(url)
	if err != nil {
		return nil, err
	}

	return checkResponse(resp)
}

// Put executes a PUT http request.
func (c *Client) Put(url string, v, body interface{}) (*Response, error) {
	resp, err := c.http.R().SetResult(v).
		SetBody(body).
		Put(url)
	if err != nil {
		return nil, err
	}

	return checkResponse(resp)
}

type Response struct {
	URL        string
	Method     string
	Status     string
	StatusCode int
	Body       string
}

func checkResponse(resp *resty.Response) (*Response, error) {
	path, _ := url.QueryUnescape(resp.Request.URL)
	r := &Response{Status: resp.Status(), StatusCode: resp.StatusCode(),
		Body: string(resp.Body()), URL: path, Method: resp.Request.Method}

	// If response is the below, return.
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return r, nil
	}

	// Otherwise, return an error
	return r, fmt.Errorf("%s %s: %d %s", r.Method, r.URL, r.StatusCode, r.Body)
}

// addQueryParams takes a slice of opts and adds each field as escaped URL query parameters to s.
// Each element in opts must be a struct whose fields contain "url" tags.
//
// Based on: https://github.com/google/go-github/blob/master/github/github.go#L226
func addQueryParams(s string, opts ...interface{}) (string, error) {
	// Handle if opts is nil
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Slice && v.IsNil() {
		return s, nil
	}

	// Parse URL
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	fulQS := url.Values{}
	for _, opt := range opts {
		qs, err := query.Values(opt)
		if err != nil {
			return s, err
		}

		for k, v := range qs {
			fulQS[k] = v
		}
	}

	u.RawQuery = fulQS.Encode()
	return u.String(), nil
}
