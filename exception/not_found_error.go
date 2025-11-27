package exception

type NotFoundError struct {
	Message string
}

func (err NotFoundError) Error() string {
	return err.Message
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		Message: message,
	}
}
