package dto

type CreateUser struct {
	User     string `json:"user"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}
type UpdateUser struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}
