package controllers

import "../Interfaces"

type AccountController struct {
	AccountManager interfaces.IAccountManager
}

// do injection of manager interface here.
func NewController(manager interfaces.IAccountManager) *AccountController {
	return &AccountController{AccountManager: manager}
}
