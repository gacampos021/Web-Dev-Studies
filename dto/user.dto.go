package dto

type CreateUser struct {
	User     string `json:"user"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
