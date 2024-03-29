package service

import (
	// "context"
	"data-entry-gamification/model"
	"data-entry-gamification/utils/errors"
	"log"

	"github.com/gin-gonic/gin"

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

	userInfo := &model.UserInfo{UserID: user.ID}
	userInfo.PostDefaultUserInfo()

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

func GetUserInfoByID(userId int64) (*model.UserInfo, *errors.RestErr) {
	result := &model.UserInfo{UserID: userId}

	if err := result.GetUserInfoByID(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetUserRolesByID(ctx *gin.Context, userId int64) ([]string, *errors.RestErr) {
	result := &model.UserInfo{UserID: userId}

	roles, err := result.UserRolesByID(ctx);
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func PutUserAvatar(ctx *gin.Context, userInfo model.UserInfo, userAvatar model.UserAvatar) (*model.UserAvatar, *errors.RestErr) {
	// TODO: validate avatar
	log.Println("service reached")
	// PUT
	if err := userInfo.UpdateAvatar(ctx, userAvatar, userInfo.UserID); err != nil {
		return nil, err
	}

	return &userAvatar, nil
}

func UserRoles(ctx *gin.Context, userId int64) ([]string, *errors.RestErr){
	result := &model.UserInfo{UserID: userId}

	userRoles, err := result.UserRolesByID(ctx)
	if err != nil {
		return nil, err
	}

	return userRoles, nil
}
