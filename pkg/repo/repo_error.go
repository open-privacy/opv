package repo

// NotFoundError ...
type NotFoundError struct {
	Err error
}

func (err NotFoundError) Error() string {
	return "NotFoundError"
}

// ValidationError ...
type ValidationError struct {
	Err     error
	Message string
}

func (err ValidationError) Error() string {
	return err.Message
}

// UnauthorizedError ...
type UnauthorizedError struct {
	Err error
}

func (err UnauthorizedError) Error() string {
	return "Unauthorized"
}
