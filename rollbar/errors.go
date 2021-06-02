package rollbar

// UserNotFoundError is returned when searching for a user by a certain criteria.
type UserNotFoundError struct {
	error
}
