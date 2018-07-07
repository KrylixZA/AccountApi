package models

type Account struct {
	AccountID int    `json:"accountId"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}
