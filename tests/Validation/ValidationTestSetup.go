package validation

import (
	validation "../../src/Validation"
)

func getSystemUnderTestValidation() *validation.Validation {
	return validation.NewValidation()
}
