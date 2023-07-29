package validators

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	// Regular expression pattern for validating email addresses
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)

	// Use the MatchString function to check if the email matches the pattern
	return re.MatchString(email)
}
