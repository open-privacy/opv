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

// NewNotFoundError ...
func NewNotFoundError(err error) error {
	return &NotFoundError{Err: err}
}

// NewValidationError ...
func NewValidationError(err error, message string) error {
	return &ValidationError{Err: err, Message: message}
}
