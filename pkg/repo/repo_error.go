package repo

type NotFoundError struct {
	Err error
}

func (err NotFoundError) Error() string {
	return "";
}

type ValidationError struct {
	Err error
	Message string
}

func (err ValidationError) Error() string {
	return err.Message;
}

func NewNotFoundError(err error) error {
	return &NotFoundError{Err:err}
}

func NewValidationError(err error, message string) error {
	return &ValidationError{Err:err, Message:message}
}

