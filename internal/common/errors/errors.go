package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeAuthorization  = ErrorType{"authorization"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
	ErrorTypeNotFound       = ErrorType{"not-found"}
)

type SlugError struct {
	error     string
	slug      string
	errorType ErrorType
}

func (s SlugError) Error() string {
	return s.error
}

func (s SlugError) Slug() string {
	return s.slug
}

func (s SlugError) ErrorType() ErrorType {
	return s.errorType
}

func NewSlugError(error, slug string) SlugError {
	return SlugError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeUnknown,
	}
}

func NewAuthorizationError(error, slug string) SlugError {
	return SlugError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeAuthorization,
	}
}

func NewIncorrectInputError(error, slug string) SlugError {
	return SlugError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeIncorrectInput,
	}
}

func NewNotFoundError(error, slug string) SlugError {
	return SlugError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeNotFound,
	}
}
