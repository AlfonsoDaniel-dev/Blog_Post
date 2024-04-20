package helpers

type Response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func NewResponse(messageType string, message string, data interface{}) Response {
	return Response{
		messageType,
		message,
		data,
	}
}
