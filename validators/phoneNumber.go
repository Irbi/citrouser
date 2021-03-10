package validators

import "regexp"

var ReNumbers *regexp.Regexp

func init() {
	ReNumbers = regexp.MustCompile(`^([+]46)\s*(7[0236])\s*(\d{4})\s*(\d{3})$`)
}

func ValidatePhoneNumber(phone string) bool {
	if !ReNumbers.MatchString(phone) {
		return false
	}
	return true
}
