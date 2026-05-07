package src

import "unicode"

func isValidPassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Minimum 8 caractères"
	}
	var lower, upper, digit, special bool
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			lower = true
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsDigit(char):
			digit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		}
	}

	if !lower || !upper || !digit || !special {
		return false, "Majuscule, minuscule, chiffre et caractère spécial requis"
	}

	return true, ""
}
