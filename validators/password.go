package validators

import "strings"

func ValidatePassword(password string) bool {

	var specialChars = "~!@#$%^&*()_+"

	if len(password) < 8 ||
		!strings.ContainsAny(password, specialChars) ||
		password == strings.ToLower(password) ||
		password == strings.ToUpper(password) {

		return false
	}

	return true
}
