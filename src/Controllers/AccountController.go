package controllers

import "../Interfaces"

type AccountController struct {
	accountManager    interfaces.IAccountManager
	requestValidation interfaces.IValidation
}

func NewController(manager interfaces.IAccountManager, validation interfaces.IValidation) *AccountController {
	return &AccountController{accountManager: manager, requestValidation: validation}
}
