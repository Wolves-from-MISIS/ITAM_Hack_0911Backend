package models

var (
	SucceedMessage = ResponseMessage{
		Message: "success",
	}
)

type ResponseMessage struct {
	Message string `json:"message"`
}
