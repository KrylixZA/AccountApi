package controllers

import "../Interfaces"

type AccountController struct {
	AccountManager interfaces.IAccountManager
}

func NewController(manager interfaces.IAccountManager) *AccountController {
	return &AccountController{AccountManager: manager}
}
