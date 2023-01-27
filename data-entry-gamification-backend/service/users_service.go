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

func GetUser(user model.User) (*model.User, *errors.RestErr) {
	result := &model.User{Email: user.Email}
	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewBadRequestError("decryption error")
	}

	resultWp := &model.User{ID: result.ID, FirstName: result.FirstName, LastName: result.LastName, Email: result.Email}
	return resultWp, nil
}

func GetUserByID(userId int64) (*model.User, *errors.RestErr) {
	result := &model.User{ID: userId}

	if err := result.GetByID(); err != nil {
		return nil, err
	}

	return result, nil
}
