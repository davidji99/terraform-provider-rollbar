// Copyright 2017
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package rollrest

// GetDateCreated returns the DateCreated field if it's non-nil, zero value otherwise.
func (i *Invitation) GetDateCreated() int64 {
	if i == nil || i.DateCreated == nil {
		return 0
	}
	return *i.DateCreated
}

// GetDateRedeemed returns the DateRedeemed field if it's non-nil, zero value otherwise.
func (i *Invitation) GetDateRedeemed() int64 {
	if i == nil || i.DateRedeemed == nil {
		return 0
	}
	return *i.DateRedeemed
}

// GetFromUserID returns the FromUserID field if it's non-nil, zero value otherwise.
func (i *Invitation) GetFromUserID() int64 {
	if i == nil || i.FromUserID == nil {
		return 0
	}
	return *i.FromUserID
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (i *Invitation) GetID() int64 {
	if i == nil || i.ID == nil {
		return 0
	}
	return *i.ID
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (i *Invitation) GetStatus() string {
	if i == nil || i.Status == nil {
		return ""
	}
	return *i.Status
}

// GetTeamID returns the TeamID field if it's non-nil, zero value otherwise.
func (i *Invitation) GetTeamID() int64 {
	if i == nil || i.TeamID == nil {
		return 0
	}
	return *i.TeamID
}

// GetToEmail returns the ToEmail field if it's non-nil, zero value otherwise.
func (i *Invitation) GetToEmail() string {
	if i == nil || i.ToEmail == nil {
		return ""
	}
	return *i.ToEmail
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (i *InvitationListResponse) GetErrorCount() int {
	if i == nil || i.ErrorCount == nil {
		return 0
	}
	return *i.ErrorCount
}

// HasResult checks if InvitationListResponse has any Result.
func (i *InvitationListResponse) HasResult() bool {
	if i == nil || i.Result == nil {
		return false
	}
	if len(i.Result) == 0 {
		return false
	}
	return true
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (i *InvitationResponse) GetErrorCount() int {
	if i == nil || i.ErrorCount == nil {
		return 0
	}
	return *i.ErrorCount
}

// GetResult returns the Result field.
func (i *InvitationResponse) GetResult() *Invitation {
	if i == nil {
		return nil
	}
	return i.Result
}

// HasScopes checks if PATCreateRequest has any Scopes.
func (p *PATCreateRequest) HasScopes() bool {
	if p == nil || p.Scopes == nil {
		return false
	}
	if len(p.Scopes) == 0 {
		return false
	}
	return true
}

// GetConfig returns the Config field.
func (p *PDRuleRequest) GetConfig() *PDRuleConfig {
	if p == nil {
		return nil
	}
	return p.Config
}

// HasFilters checks if PDRuleRequest has any Filters.
func (p *PDRuleRequest) HasFilters() bool {
	if p == nil || p.Filters == nil {
		return false
	}
	if len(p.Filters) == 0 {
		return false
	}
	return true
}

// GetAccountID returns the AccountID field if it's non-nil, zero value otherwise.
func (p *Project) GetAccountID() int64 {
	if p == nil || p.AccountID == nil {
		return 0
	}
	return *p.AccountID
}

// GetDataCreated returns the DataCreated field if it's non-nil, zero value otherwise.
func (p *Project) GetDataCreated() int64 {
	if p == nil || p.DataCreated == nil {
		return 0
	}
	return *p.DataCreated
}

// GetDateModified returns the DateModified field if it's non-nil, zero value otherwise.
func (p *Project) GetDateModified() int64 {
	if p == nil || p.DateModified == nil {
		return 0
	}
	return *p.DateModified
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (p *Project) GetID() int64 {
	if p == nil || p.ID == nil {
		return 0
	}
	return *p.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (p *Project) GetName() string {
	if p == nil || p.Name == nil {
		return ""
	}
	return *p.Name
}

// GetSettingsData returns the SettingsData field.
func (p *Project) GetSettingsData() *ProjectSD {
	if p == nil {
		return nil
	}
	return p.SettingsData
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (p *Project) GetStatus() string {
	if p == nil || p.Status == nil {
		return ""
	}
	return *p.Status
}

// GetAccessToken returns the AccessToken field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetAccessToken() string {
	if p == nil || p.AccessToken == nil {
		return ""
	}
	return *p.AccessToken
}

// GetCurrentRateLimitWindowCount returns the CurrentRateLimitWindowCount field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetCurrentRateLimitWindowCount() int64 {
	if p == nil || p.CurrentRateLimitWindowCount == nil {
		return 0
	}
	return *p.CurrentRateLimitWindowCount
}

// GetCurrentRateLimitWindowStart returns the CurrentRateLimitWindowStart field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetCurrentRateLimitWindowStart() int64 {
	if p == nil || p.CurrentRateLimitWindowStart == nil {
		return 0
	}
	return *p.CurrentRateLimitWindowStart
}

// GetDataCreated returns the DataCreated field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetDataCreated() int64 {
	if p == nil || p.DataCreated == nil {
		return 0
	}
	return *p.DataCreated
}

// GetDateModified returns the DateModified field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetDateModified() int64 {
	if p == nil || p.DateModified == nil {
		return 0
	}
	return *p.DateModified
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetName() string {
	if p == nil || p.Name == nil {
		return ""
	}
	return *p.Name
}

// GetProjectID returns the ProjectID field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetProjectID() int64 {
	if p == nil || p.ProjectID == nil {
		return 0
	}
	return *p.ProjectID
}

// GetRateLimitWindowCount returns the RateLimitWindowCount field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetRateLimitWindowCount() int {
	if p == nil || p.RateLimitWindowCount == nil {
		return 0
	}
	return *p.RateLimitWindowCount
}

// GetRateLimitWindowSize returns the RateLimitWindowSize field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetRateLimitWindowSize() int {
	if p == nil || p.RateLimitWindowSize == nil {
		return 0
	}
	return *p.RateLimitWindowSize
}

// HasScopes checks if ProjectAccessToken has any Scopes.
func (p *ProjectAccessToken) HasScopes() bool {
	if p == nil || p.Scopes == nil {
		return false
	}
	if len(p.Scopes) == 0 {
		return false
	}
	return true
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (p *ProjectAccessToken) GetStatus() string {
	if p == nil || p.Status == nil {
		return ""
	}
	return *p.Status
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (p *ProjectAccessTokenListResponse) GetErrorCount() int {
	if p == nil || p.ErrorCount == nil {
		return 0
	}
	return *p.ErrorCount
}

// HasResult checks if ProjectAccessTokenListResponse has any Result.
func (p *ProjectAccessTokenListResponse) HasResult() bool {
	if p == nil || p.Result == nil {
		return false
	}
	if len(p.Result) == 0 {
		return false
	}
	return true
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (p *ProjectAccessTokenResponse) GetErrorCount() int {
	if p == nil || p.ErrorCount == nil {
		return 0
	}
	return *p.ErrorCount
}

// GetResult returns the Result field.
func (p *ProjectAccessTokenResponse) GetResult() *ProjectAccessToken {
	if p == nil {
		return nil
	}
	return p.Result
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (p *ProjectListResponse) GetErrorCount() int {
	if p == nil || p.ErrorCount == nil {
		return 0
	}
	return *p.ErrorCount
}

// HasResult checks if ProjectListResponse has any Result.
func (p *ProjectListResponse) HasResult() bool {
	if p == nil || p.Result == nil {
		return false
	}
	if len(p.Result) == 0 {
		return false
	}
	return true
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (p *ProjectResponse) GetErrorCount() int {
	if p == nil || p.ErrorCount == nil {
		return 0
	}
	return *p.ErrorCount
}

// GetResult returns the Result field.
func (p *ProjectResponse) GetResult() *Project {
	if p == nil {
		return nil
	}
	return p.Result
}

// GetFingerprintVersions returns the FingerprintVersions field.
func (p *ProjectSD) GetFingerprintVersions() *ProjectSDFingerprintVersions {
	if p == nil {
		return nil
	}
	return p.FingerprintVersions
}

// GetMigrations returns the Migrations field.
func (p *ProjectSD) GetMigrations() *ProjectSDMigrations {
	if p == nil {
		return nil
	}
	return p.Migrations
}

// GetAndroidAndroid returns the AndroidAndroid field if it's non-nil, zero value otherwise.
func (p *ProjectSDFingerprintVersions) GetAndroidAndroid() int {
	if p == nil || p.AndroidAndroid == nil {
		return 0
	}
	return *p.AndroidAndroid
}

// GetBrowserBrowserJS returns the BrowserBrowserJS field if it's non-nil, zero value otherwise.
func (p *ProjectSDFingerprintVersions) GetBrowserBrowserJS() int {
	if p == nil || p.BrowserBrowserJS == nil {
		return 0
	}
	return *p.BrowserBrowserJS
}

// GetUnminifyReactErrors returns the UnminifyReactErrors field if it's non-nil, zero value otherwise.
func (p *ProjectSDFingerprintVersions) GetUnminifyReactErrors() int {
	if p == nil || p.UnminifyReactErrors == nil {
		return 0
	}
	return *p.UnminifyReactErrors
}

// GetEnableCalculateSymbolRanges returns the EnableCalculateSymbolRanges field if it's non-nil, zero value otherwise.
func (p *ProjectSDMigrations) GetEnableCalculateSymbolRanges() int {
	if p == nil || p.EnableCalculateSymbolRanges == nil {
		return 0
	}
	return *p.EnableCalculateSymbolRanges
}

// GetEnableCustomFingerprintingOverride returns the EnableCustomFingerprintingOverride field if it's non-nil, zero value otherwise.
func (p *ProjectSDMigrations) GetEnableCustomFingerprintingOverride() int {
	if p == nil || p.EnableCustomFingerprintingOverride == nil {
		return 0
	}
	return *p.EnableCustomFingerprintingOverride
}

// GetEnableMissingJquery returns the EnableMissingJquery field if it's non-nil, zero value otherwise.
func (p *ProjectSDMigrations) GetEnableMissingJquery() int {
	if p == nil || p.EnableMissingJquery == nil {
		return 0
	}
	return *p.EnableMissingJquery
}

// GetEnableSourceMaps returns the EnableSourceMaps field if it's non-nil, zero value otherwise.
func (p *ProjectSDMigrations) GetEnableSourceMaps() int {
	if p == nil || p.EnableSourceMaps == nil {
		return 0
	}
	return *p.EnableSourceMaps
}

// GetRecognizeDirectRecursion returns the RecognizeDirectRecursion field if it's non-nil, zero value otherwise.
func (p *ProjectSDMigrations) GetRecognizeDirectRecursion() int {
	if p == nil || p.RecognizeDirectRecursion == nil {
		return 0
	}
	return *p.RecognizeDirectRecursion
}

// GetAccessLevel returns the AccessLevel field if it's non-nil, zero value otherwise.
func (t *Team) GetAccessLevel() string {
	if t == nil || t.AccessLevel == nil {
		return ""
	}
	return *t.AccessLevel
}

// GetAccountID returns the AccountID field if it's non-nil, zero value otherwise.
func (t *Team) GetAccountID() int64 {
	if t == nil || t.AccountID == nil {
		return 0
	}
	return *t.AccountID
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (t *Team) GetID() int64 {
	if t == nil || t.ID == nil {
		return 0
	}
	return *t.ID
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (t *Team) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (t *TeamListResponse) GetErrorCount() int {
	if t == nil || t.ErrorCount == nil {
		return 0
	}
	return *t.ErrorCount
}

// HasResult checks if TeamListResponse has any Result.
func (t *TeamListResponse) HasResult() bool {
	if t == nil || t.Result == nil {
		return false
	}
	if len(t.Result) == 0 {
		return false
	}
	return true
}

// GetProjectID returns the ProjectID field if it's non-nil, zero value otherwise.
func (t *TeamProjectAssoc) GetProjectID() int64 {
	if t == nil || t.ProjectID == nil {
		return 0
	}
	return *t.ProjectID
}

// GetTeamID returns the TeamID field if it's non-nil, zero value otherwise.
func (t *TeamProjectAssoc) GetTeamID() int64 {
	if t == nil || t.TeamID == nil {
		return 0
	}
	return *t.TeamID
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (t *TeamProjectAssocListResponse) GetErrorCount() int {
	if t == nil || t.ErrorCount == nil {
		return 0
	}
	return *t.ErrorCount
}

// HasResult checks if TeamProjectAssocListResponse has any Result.
func (t *TeamProjectAssocListResponse) HasResult() bool {
	if t == nil || t.Result == nil {
		return false
	}
	if len(t.Result) == 0 {
		return false
	}
	return true
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (t *TeamProjectAssocResponse) GetErrorCount() int {
	if t == nil || t.ErrorCount == nil {
		return 0
	}
	return *t.ErrorCount
}

// GetResult returns the Result field.
func (t *TeamProjectAssocResponse) GetResult() *TeamProjectAssoc {
	if t == nil {
		return nil
	}
	return t.Result
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (t *TeamResponse) GetErrorCount() int {
	if t == nil || t.ErrorCount == nil {
		return 0
	}
	return *t.ErrorCount
}

// GetResult returns the Result field.
func (t *TeamResponse) GetResult() *Team {
	if t == nil {
		return nil
	}
	return t.Result
}

// GetTeamID returns the TeamID field if it's non-nil, zero value otherwise.
func (t *TeamUserAssoc) GetTeamID() int64 {
	if t == nil || t.TeamID == nil {
		return 0
	}
	return *t.TeamID
}

// GetUserID returns the UserID field if it's non-nil, zero value otherwise.
func (t *TeamUserAssoc) GetUserID() int64 {
	if t == nil || t.UserID == nil {
		return 0
	}
	return *t.UserID
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (t *TeamUserListResponse) GetErrorCount() int {
	if t == nil || t.ErrorCount == nil {
		return 0
	}
	return *t.ErrorCount
}

// HasResult checks if TeamUserListResponse has any Result.
func (t *TeamUserListResponse) HasResult() bool {
	if t == nil || t.Result == nil {
		return false
	}
	if len(t.Result) == 0 {
		return false
	}
	return true
}

// GetEmail returns the Email field if it's non-nil, zero value otherwise.
func (u *User) GetEmail() string {
	if u == nil || u.Email == nil {
		return ""
	}
	return *u.Email
}

// GetEmailEnabled returns the EmailEnabled field if it's non-nil, zero value otherwise.
func (u *User) GetEmailEnabled() int {
	if u == nil || u.EmailEnabled == nil {
		return 0
	}
	return *u.EmailEnabled
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (u *User) GetID() int64 {
	if u == nil || u.ID == nil {
		return 0
	}
	return *u.ID
}

// GetUsername returns the Username field if it's non-nil, zero value otherwise.
func (u *User) GetUsername() string {
	if u == nil || u.Username == nil {
		return ""
	}
	return *u.Username
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (u *UserListResponse) GetErrorCount() int {
	if u == nil || u.ErrorCount == nil {
		return 0
	}
	return *u.ErrorCount
}

// GetResult returns the Result field.
func (u *UserListResponse) GetResult() *UserListResult {
	if u == nil {
		return nil
	}
	return u.Result
}

// HasUsers checks if UserListResult has any Users.
func (u *UserListResult) HasUsers() bool {
	if u == nil || u.Users == nil {
		return false
	}
	if len(u.Users) == 0 {
		return false
	}
	return true
}

// GetAccountID returns the AccountID field if it's non-nil, zero value otherwise.
func (u *UserProject) GetAccountID() int64 {
	if u == nil || u.AccountID == nil {
		return 0
	}
	return *u.AccountID
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (u *UserProject) GetID() int64 {
	if u == nil || u.ID == nil {
		return 0
	}
	return *u.ID
}

// GetSlug returns the Slug field if it's non-nil, zero value otherwise.
func (u *UserProject) GetSlug() string {
	if u == nil || u.Slug == nil {
		return ""
	}
	return *u.Slug
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (u *UserProject) GetStatus() int {
	if u == nil || u.Status == nil {
		return 0
	}
	return *u.Status
}

// HasProjects checks if UserProjectsList has any Projects.
func (u *UserProjectsList) HasProjects() bool {
	if u == nil || u.Projects == nil {
		return false
	}
	if len(u.Projects) == 0 {
		return false
	}
	return true
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (u *UserProjectsListResponse) GetErrorCount() int {
	if u == nil || u.ErrorCount == nil {
		return 0
	}
	return *u.ErrorCount
}

// GetResult returns the Result field.
func (u *UserProjectsListResponse) GetResult() *UserProjectsList {
	if u == nil {
		return nil
	}
	return u.Result
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (u *UserResponse) GetErrorCount() int {
	if u == nil || u.ErrorCount == nil {
		return 0
	}
	return *u.ErrorCount
}

// GetResult returns the Result field.
func (u *UserResponse) GetResult() *User {
	if u == nil {
		return nil
	}
	return u.Result
}

// HasTeams checks if UserTeamsList has any Teams.
func (u *UserTeamsList) HasTeams() bool {
	if u == nil || u.Teams == nil {
		return false
	}
	if len(u.Teams) == 0 {
		return false
	}
	return true
}

// GetErrorCount returns the ErrorCount field if it's non-nil, zero value otherwise.
func (u *UserTeamsListResponse) GetErrorCount() int {
	if u == nil || u.ErrorCount == nil {
		return 0
	}
	return *u.ErrorCount
}

// GetResult returns the Result field.
func (u *UserTeamsListResponse) GetResult() *UserTeamsList {
	if u == nil {
		return nil
	}
	return u.Result
}
