package diagnostics

func GetErrorDescription(code int) string {
	switch code {
	case BadRequest:
		return "The request was malformed, nil or empty."
	case BadLogin:
		return "The login supplied in the request body was nil, empty or malformed."
	case BadEmail:
		return "The email supplied in the request body was nil, empty, malformed or did not contain and '@'."
	case BadFirstName:
		return "The first name supplied in the request body was nil, empty or malformed."
	case BadSurname:
		return "The surname supplied in the request body was nil, empty or malformed."
	case BadPassword:
		return "The password supplied in the request body was nil, empty or malformed."
	case BadUsername:
		return "The username supplied in the request body was nil, empty or malformed."
	case BadCurrentPassword:
		return "The password supplied as the 'current password' in the request body was nil, empty or malformed."
	case BadNewPassword:
		return "The password supplied as the 'new password' in the request body was nil, empty or malformed."
	case FailedToCreateAccount:
		return "There was an error creating the account on the database. Please try again later."
	case NoAccountDetailsFound:
		return "No accounts were found for the given account identifier."
	case LoginFailed:
		return "Login failed as no account was found that matched the given login and password."
	case ResetPasswordFailed:
		return "Failed to reset account password as no account could be found for the given account identifier with the given password."
	}

	return "An unknown error occurred."
}
