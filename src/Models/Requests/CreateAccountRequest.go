package requests

type CreateAccountRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Password  string `json:"password"`
}
