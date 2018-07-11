package controllers

import (
	controllers "../../src/Controllers"
	interfaces "../../src/Interfaces"
)

func getSystemUnderTestAccountController(manager interfaces.IAccountManager) *controllers.AccountController {
	return controllers.NewController(manager)
}
