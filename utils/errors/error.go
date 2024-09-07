package errors

type ErrorMeta struct {
	Message string `json:"message"` // Message 错误信息
}

func (e *ErrorMeta) Error() string {
	return e.Message
}

func newErrorMeta(message string) *ErrorMeta {
	return &ErrorMeta{
		Message: message,
	}
}

type BadRequestError struct {
	*ErrorMeta
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		ErrorMeta: newErrorMeta(message),
	}
}

type NotFoundError struct {
	*ErrorMeta
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		ErrorMeta: newErrorMeta(message),
	}
}

type UnauthorizedError struct {
	*ErrorMeta
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		ErrorMeta: newErrorMeta(message),
	}
}
