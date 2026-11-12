package api

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Errors interface {
	StatusCode() int
	Err() []Error
	Fail() bool
	Error() string
}

type Error interface {
	StatusCode() int
	Errors() []Error
	Err() Error
	Fail() bool
	Error() string
}

type ValidationErrors struct {
	errors []Error
}

func (ve *ValidationErrors) StatusCode() int {
	return 400
}
func (ve *ValidationErrors) Fail() bool {
	return len(ve.errors) > 0
}
func (ve *ValidationErrors) Errors() []Error {
	return ve.errors
}
func (ve *ValidationErrors) AddError(err Error) {
	ve.errors = append(ve.errors, err)
}
func (ve *ValidationErrors) Err() Error {
	return ve
}
func (ve *ValidationErrors) Error() string {
	return "Validation errors occurred"
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	fail    bool
	code    int
}

func (ve *ValidationError) StatusCode() int {
	return ve.code
}

func (ve *ValidationError) Err() Error {
	return ve
}

func (ve *ValidationError) Error() string {
	return ve.Message
}

func (ve *ValidationError) Errors() []Error {
	return []Error{ve}
}

func (ve *ValidationError) Fail() bool {
	return ve.fail
}

func CreateError(field string, message string) Error {
	return &ValidationError{
		Field:   field,
		Message: message,
		fail:    false,
		code:    400,
	}
}

func FromError(title string, e error) Error {
	return &ValidationError{
		Message: e.Error(),
	}
}
