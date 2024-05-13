package responses

type Response struct {
	Data    interface{} `json:"data"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
}

const Message = "OK"
const Error = "Error"

func NewResponse(data interface{}, message string, code string) Response {
	return Response{
		Data:    data,
		Code:    code,
		Message: message,
	}
}
