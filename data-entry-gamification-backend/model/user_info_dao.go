package model

import (
	"data-entry-gamification/storage/users_db"
	"data-entry-gamification/utils/errors"
)

var (
	queryGetUserInfoByID = "SELECT user_id, points, level FROM user_info WHERE user_id=?;"
)

func (userInfo *UserInfo) GetUserInfoByID() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserInfoByID)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(userInfo.UserID)
	if getErr := result.Scan(&userInfo.UserID, &userInfo.Points, &userInfo.Level); getErr != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}