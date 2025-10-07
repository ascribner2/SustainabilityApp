package services

import (
	"regexp"

	"github.com/ascribner/sustainabilityapp/internal/entity"
)

func TotalItemOffset(items []entity.Item) float64 {
	var total float64 = 0.0

	for _, item := range items {
		total += item.GetOffset()
	}

	return total
}

func VerifyEmail(email string) bool {
	re := regexp.MustCompile(`^[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}$`)

	return re.MatchString(email)
}

func VerifyPassword(password string) bool {
	// check length
	if len(password) < 6 {
		return false
	}

	// check special character
	if valid, err := regexp.MatchString(`[!@#$%^&*]`, password); !valid || err != nil {
		return false
	}

	// check number
	if valid, err := regexp.MatchString(`[0-9]`, password); !valid || err != nil {
		return false
	}

	// valid password
	return true
}
