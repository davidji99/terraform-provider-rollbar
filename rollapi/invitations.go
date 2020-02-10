package rollapi

// InvitationService handles communication with the invitation related
// methods of the Rollbar API.
//
// Rollbar API docs: https://docs.rollbar.com/invitation
type InvitationService service

// InvitationResponse represents a response after inviting an user.
type InvitationResponse struct {
	ErrorCount *int        `json:"err,omitempty"`
	Result     *Invitation `json:"result,omitempty"`
}

// InvitationListResponse represents a response of all invitations.
type InvitationListResponse struct {
	ErrorCount *int          `json:"err,omitempty"`
	Result     []*Invitation `json:"result,omitempty"`
}

// Invitation represents an invitation in Rollbar (usually an user's invitation to a team).
type Invitation struct {
	ID           *int64  `json:"id,omitempty"`
	FromUserID   *int64  `json:"from_user_id,omitempty"`
	TeamID       *int64  `json:"team_id,omitempty"`
	ToEmail      *string `json:"to_email,omitempty"`
	Status       *string `json:"status,omitempty"`
	DateCreated  *int64  `json:"date_created,omitempty"`
	DateRedeemed *int64  `json:"date_redeemed,omitempty"`
}

// Get a invitation.
//
// Rollbar API docs: https://docs.rollbar.com/reference#get-invitation
func (i *InvitationService) Get(inviteID int) (*InvitationResponse, *Response, error) {
	var result *InvitationResponse
	urlStr := i.client.requestURL("/invite/%d", inviteID)

	// Set the correct authentication header
	i.client.setAuthTokenHeader(i.client.accountAccessToken)

	// Execute the request
	response, getErr := i.client.Get(urlStr, &result, nil)

	return result, response, getErr
}

// Cancel a invitation.
//
// Rollbar API docs: https://docs.rollbar.com/reference#cancel-invitation
func (i *InvitationService) Cancel(inviteID int) (*Response, error) {
	urlStr := i.client.requestURL("/invite/%d", inviteID)

	// Set the correct authentication header
	i.client.setAuthTokenHeader(i.client.accountAccessToken)

	// Execute the request
	response, getErr := i.client.Delete(urlStr, nil)

	return response, getErr
}
