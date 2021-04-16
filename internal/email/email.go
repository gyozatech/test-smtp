package email

import (
	"errors"
	"regexp"
	"strings"
)

// Validate if email address is formally correct
func Validate(addr string) bool {
	emailValidator := regexp.MustCompile(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]{2,}`)
	valid := emailValidator.MatchString(addr)
	return valid
}

// Split extracts local name and host from a given email address
func Split(addr string) (local, domain string, err error) {
	if !Validate(addr) {
		return "", "", errors.New("mta: invalid mail address")
	}
	parts := strings.SplitN(addr, "@", 2)

	if len(parts) != 2 {
		// Should never be called!
		return "", "", errors.New("mta: invalid mail address")
	}
	return parts[0], parts[1], nil
}
