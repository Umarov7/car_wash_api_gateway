package middleware

import (
	"api-gateway/config"
	pbu "api-gateway/genproto/user"
	"api-gateway/pkg"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func ValidateUser(cfg *config.Config, userID string) error {
	_, err := uuid.Parse(userID)
	if err != nil {
		return errors.Wrap(err, "invalid user id")
	}

	user := pkg.NewUserClient(cfg)

	_, err = user.ValidateUser(context.Background(), &pbu.ID{Id: userID})
	if err != nil {
		return errors.Wrap(err, "user not found")
	}

	return nil
}
