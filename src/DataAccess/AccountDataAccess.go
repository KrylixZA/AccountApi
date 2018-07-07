package dataaccess

import (
	models "../Models"
)

var accounts map[int]models.Account

func setupAccounts() {
	accounts = make(map[int]models.Account)
	accounts[1] = models.Account{AccountID: 1, Login: "headleysj@gmail.com", Password: "password1234", FirstName: "Simon", Surname: "Headley", Email: "headleysj@gmail.com"}
}

type AccountDataAccess struct {
}
