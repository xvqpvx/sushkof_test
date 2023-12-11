package request

type SaveUserRequest struct {
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
