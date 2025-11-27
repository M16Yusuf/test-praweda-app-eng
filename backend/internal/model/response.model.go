package model

type Response struct {
	IsSuccess bool   `json:"is_success"  example:"true"`
	Code      int    `json:"code,omitempty"  example:"200"`
	Page      int    `json:"page,omitempty"  example:"1"`
	Msg       string `json:"message,omitempty"  example:"Example message success..."`
}

type ErrorResponse struct {
	Response
	Err string `json:"error" example:"Error message..."`
}

type ResponseData struct {
	Response
	Data interface{} `json:"data,omitempty"`
}
