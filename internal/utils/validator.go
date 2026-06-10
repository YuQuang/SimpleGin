package utils

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	var emailRegex = regexp.MustCompile(
		`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
	)
	if !emailRegex.MatchString(email) {
		return false
	}

	return true
}

func ValidateUsername(username string) bool {
	hasSpecial, _ := regexp.MatchString(`[!@#$%^&*()\-+=\[\]{};:'",<>/?\\|~]`, username)
	if hasSpecial {
		return false
	}

	if 3 > len(username) || len(username) > 255 {
		return false
	}

	return true
}

func ValidatePassword(password string) bool {
	hasUpper, _ := regexp.MatchString(`[A-Z]`, password)
	if !hasUpper {
		return false
	}

	hasNumber, _ := regexp.MatchString(`[0-9]`, password)
	if !hasNumber {
		return false
	}

	hasSpecial, _ := regexp.MatchString(`[^A-Za-z0-9]`, password)
	if !hasSpecial {
		return false
	}

	if 10 > len(password) || len(password) > 255 {
		return false
	}

	return true
}
