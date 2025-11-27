package model

type Response struct {
	IsSuccess bool   `json:"is_success"  example:"true"`
	Code      int    `json:"code,omitempty"  example:"200"`
	Page      int    `json:"page,omitempty"  example:"1"`
	Err       string `json:"error,omitempty" example:"Error message..."`
	Msg       string `json:"message,omitempty"  example:"Example message success..."`
}

type ResponseData struct {
	Response
	Data interface{} `json:"data,omitempty"`
}
