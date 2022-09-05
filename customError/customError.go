package customerror

type CustomError struct {
	Message    any
	StatusCode int
}

func NewError(message any, code int) *CustomError {
	return &CustomError{
		message, code,
	}
}

var (
	UnSupportedFileType = &CustomError{
		Message:    "Only pdf, jpeg, jpg and png files are supported",
		StatusCode: 400,
	}
	UnForbiddenError = &CustomError{
		Message:    "Unforbidden access",
		StatusCode: 403,
	}
	InvalidEntity = &CustomError{
		Message:    "Unsupported entity found at database",
		StatusCode: 500,
	}
	NotFoundError = &CustomError{
		Message:    "Document not found",
		StatusCode: 404,
	}
)
