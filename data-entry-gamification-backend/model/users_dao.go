package model

import (
	"data-entry-gamification/storage/users_db"
	"data-entry-gamification/utils/errors"
)

var (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	user.ID = userID
	return nil
}
