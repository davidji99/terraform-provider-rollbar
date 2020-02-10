package rollapi

// UsersService handles communication with the users related
// methods of the Rollbar API.
//
// Rollbar API docs: https://docs.rollbar.com/reference#users
type UsersService service

// User represents a user in Rollbar.
type User struct {
	ID       *int64  `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
}

// UserResponse represents the response returned after getting a user.
type UserResponse struct {
	ErrorCount int   `json:"err,omitempty"`
	Result     *User `json:"result,omitempty"`
}

// UserListResponse represents the response returned after getting all users.
type UserListResponse struct {
	ErrorCount int     `json:"err,omitempty"`
	Result     []*User `json:"result,omitempty"`
}

// List all users.
//
// Rollbar API docs: https://docs.rollbar.com/reference#list-all-users
func (u *UsersService) List() (*UserListResponse, *Response, error) {
	var result *UserListResponse
	urlStr := u.client.requestURL("/users")

	// Set the correct authentication header
	u.client.setAuthTokenHeader(u.client.accountAccessToken)

	// Execute the request
	response, getErr := u.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Get a users.
//
// Returns basic information about the user, as relevant to the account your access token is for.
// This is the same information available on the "Members" page in the Rollbar UI.
//
// Rollbar API docs: https://docs.rollbar.com/reference#get-a-user
func (u *UsersService) Get(userId int) (*UserResponse, *Response, error) {
	var result *UserResponse
	urlStr := u.client.requestURL("/user/%d", userId)

	// Set the correct authentication header
	u.client.setAuthTokenHeader(u.client.accountAccessToken)

	// Execute the request
	response, getErr := u.client.Get(urlStr, &result, nil)

	return result, response, getErr
}
