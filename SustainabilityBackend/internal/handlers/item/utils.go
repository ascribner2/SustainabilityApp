package item

import (
	"context"
	"errors"

	"github.com/ascribner/sustainabilityapp/internal/contextkeys"
)

type AuthContextKey string

func GetEmailFromContext(ctx context.Context) (string, error) {
	var emailKey contextkeys.AuthContextKey = "email"

	email, ok := ctx.Value(emailKey).(string)
	if !ok {
		return "", errors.New("underlying type not correct")
	}

	return email, nil
}
