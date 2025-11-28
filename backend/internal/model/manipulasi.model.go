package model

type ManipulatedData struct {
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Email    string   `json:"email"`
	Age      int      `json:"age"`
	Phone    string   `json:"phone"`
	Cell     string   `json:"cell"`
	Picture  []string `json:"picture"`
}
