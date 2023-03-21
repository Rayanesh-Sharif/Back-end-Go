package util

import "net/mail"

// CheckEmail will check if a given value to email is valid or not.
// If it is, it will return the email. Otherwise, an empty string is returned.
func CheckEmail(email string) string {
	m, err := mail.ParseAddress(email)
	if err != nil {
		return ""
	}
	return m.Address
}
