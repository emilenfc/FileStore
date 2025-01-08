package utils

import (
	"crypto/rand"
	"encoding/base64"
	"regexp"
	"strings"
)

func GenerateAPIKey() string {
	key := make([]byte, 7)
	rand.Read(key)
	return strings.ToUpper(base64.URLEncoding.EncodeToString(key))[:10]

}

func GenerateAPISecret() string {
	key := make([]byte, 4)
	rand.Read(key)
	return base64.URLEncoding.EncodeToString(key)[:5]
}

func IsValidPhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`)
	return phoneRegex.MatchString(phone)
}
