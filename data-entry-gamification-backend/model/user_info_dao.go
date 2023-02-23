package model

import (
	"data-entry-gamification/storage/users_db"
	"data-entry-gamification/utils/errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	queryGetUserInfoByID = "SELECT user_id, points, level, img_uri FROM user_info WHERE user_id=?;"
	queryUpdateUserImgURI = "UPDATE user_info SET img_uri=? WHERE user_id=?;"
	queryGetUserRolesByID = "SELECT user_id, user_role FROM user_roles WHERE user_id=?;"
	queryInsertDefaultUserInfo = "INSERT INTO user_info (user_id, points, level, img_uri) VALUES (?, 0, 0, '../assets/user-avatars/001-Default-Avatar-2.jpg');"
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

func (userInfo *UserInfo) UserRolesByID(ctx *gin.Context) ([]string, *errors.RestErr) {
	log.Println("user ID is", userInfo.UserID)
	stmt, err := users_db.Client.Prepare(queryGetUserRolesByID)
	if err != nil {
		return nil, errors.NewInternalServerError("database get user roles error")
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userInfo.UserID)
	var allUserRoles []string;
	for rows.Next() {
		roleEntry := new(UserRole)
		if getErr := rows.Scan(&roleEntry.UserID, &roleEntry.UserRole); getErr != nil {
			return nil, errors.NewInternalServerError("database user role extraction error")
		}
		allUserRoles = append(allUserRoles, roleEntry.UserRole)
	}

	return allUserRoles, nil
}

func (userInfo *UserInfo) PostDefaultUserInfo() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertDefaultUserInfo)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()	
	insertResult, saveErr := stmt.Exec(userInfo.UserID)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}
	log.Println("insertResult", insertResult)
	return nil
}