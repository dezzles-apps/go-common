package responses

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

func (e *Error) Error() string {
	return e.Message
}

func CreateError(field string, message string) Error {
	var err Error
	err.Field = field
	err.Message = message
	return err
}
