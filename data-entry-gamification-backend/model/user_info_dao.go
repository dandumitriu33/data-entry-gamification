package model

import (
	// "context"
	"data-entry-gamification/storage/users_db"
	"data-entry-gamification/utils/errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	queryGetUserInfoByID = "SELECT user_id, points, level, img_uri FROM user_info WHERE user_id=?;"
	queryUpdateUserImgURI = "UPDATE user_info SET img_uri=? WHERE user_id=?;"
)

func (userInfo *UserInfo) GetUserInfoByID() *errors.RestErr {
	log.Println("user ID is", userInfo.UserID)
	stmt, err := users_db.Client.Prepare(queryGetUserInfoByID)
	if err != nil {
		return errors.NewInternalServerError("database get user info error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(userInfo.UserID)
	if getErr := result.Scan(&userInfo.UserID, &userInfo.Points, &userInfo.Level, &userInfo.ImageURI); getErr != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (userInfo *UserInfo) UpdateAvatar(ctx *gin.Context, userAvatar UserAvatar, userID int64) *errors.RestErr {
	log.Println("upload transaction starting")
	// Begin GIN transaction
	txHandle, err := users_db.Client.Begin()
	if err != nil {
		return errors.NewInternalServerError("database transaction error")
	}
	defer txHandle.Rollback()
	log.Println("transaction started")

	// Upload the image
	currentTime := time.Now()
	date_added := currentTime.Format("20060102150405")
	filePathOnServer := "../assets/user-avatars/"+date_added+"_"+userAvatar.Avatar.Filename
	err = ctx.SaveUploadedFile(userAvatar.Avatar, filePathOnServer)
	if err != nil {
		log.Println(err.Error())
		return errors.NewInternalServerError("file upload failed")
	}
	log.Println("file uploaded")

	// Update the Image URI in DB on User Info
	_, pairErr := txHandle.ExecContext(ctx, queryUpdateUserImgURI, filePathOnServer, userID)
	if pairErr != nil {
		return errors.NewInternalServerError("database transaction pair error")
	}
	log.Println("uri updated")

	commitErr := txHandle.Commit(); 
	if commitErr != nil {
		return errors.NewInternalServerError("database transaction commit error")
	}
	log.Println("transaction commited")
	return nil
}
