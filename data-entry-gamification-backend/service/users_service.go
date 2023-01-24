package service

import (
	"data-entry-gamification/model"
	"data-entry-gamification/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user model.User) (*model.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// encrypt the password
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt the pasword")
	}

	user.Password = string(pwSlice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
