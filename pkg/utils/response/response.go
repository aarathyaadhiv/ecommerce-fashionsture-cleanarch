package response

type Response struct {
	Statuscode uint
	Message    string
	Data       interface{}
	Error      interface{}
}

func Responses(code uint, message string, data interface{}, error interface{}) Response {
	return Response{Statuscode: code,
		Message: message,
		Data:    data,
		Error:   error,
	}
}
