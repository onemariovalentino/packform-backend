package helper

type (
	NewResponseData struct {
		RequestID    string      `json:"request_id"`
		Code         int         `json:"code"`
		Message      string      `json:"message"`
		ErrorMessage string      `json:"error_message"`
		Data         interface{} `json:"data"`
	}
)

func NewResponse(code int, message, error_message string, data interface{}) *NewResponseData {
	return &NewResponseData{Code: code, Message: message, ErrorMessage: error_message, Data: data}
}
