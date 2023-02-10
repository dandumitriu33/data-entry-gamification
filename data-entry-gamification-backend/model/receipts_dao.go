package model

import (
	"context"
	"data-entry-gamification/storage/receipts_db"
	"data-entry-gamification/utils/errors"
	"time"
)

var (
	queryInsertReceipt         = "INSERT INTO receipts (model_year, make, vin, first_name, last_name, state, date_added) VALUES (?, ?, ?, ?, ?, ?, ?);"
	queryGetAllCount           = "SELECT COUNT(*) FROM receipts;"
	queryGetAllCountToday      = "SELECT COUNT(*) FROM receipts WHERE DATE(date_added) = DATE(NOW());"
	queryInsertUserIDReceiptID = "INSERT INTO user_receipts (user_id, receipt_id) VALUES (?, ?);"
	queryGetUserPointsByUserID = "SELECT points FROM user_info WHERE user_id = ?;"
	queryInsertNewUserPoints   = "INSERT INTO user_info (user_id, points, level) VALUES (?, ?, ?);"
	queryUpdateUserPoints      = "UPDATE user_info SET points = ?, level = ? WHERE user_id = ?;"
)

func (receipt *Receipt) Save(ctx context.Context, userID int64) *errors.RestErr {
	tx, err := receipts_db.Client.BeginTx(ctx, nil)
	if err != nil {
		return errors.NewInternalServerError("database transaction error")
	}
	defer tx.Rollback()

	// Add receipt and get the ID
	currentTime := time.Now()
	date_added := currentTime.Format("20060102150405")
	addResult, addErr := tx.ExecContext(ctx, queryInsertReceipt, receipt.ModelYear, receipt.Make, receipt.Vin, receipt.FirstName, receipt.LastName, receipt.State, date_added)
	if addErr != nil {
		return errors.NewInternalServerError("database transaction add error")
	}
	receiptID, err := addResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database transaction receipt ID retrieval error")
	}
	receipt.ID = receiptID

	// Add userID and receiptID to user_receipts table
	_, pairErr := tx.ExecContext(ctx, queryInsertUserIDReceiptID, userID, receiptID)
	if pairErr != nil {
		return errors.NewInternalServerError("database transaction pair error")
	}

	// Get User points
	var userPoints int64
	getPointsErr := tx.QueryRowContext(ctx, queryGetUserPointsByUserID, userID).Scan(&userPoints)
	if getPointsErr != nil {
		return errors.NewInternalServerError("database transaction get points error")
	}

	if userPoints == 0 {
		// Insert User Points - first user input recording
		_, pointsErr := tx.ExecContext(ctx, queryInsertNewUserPoints, userID, 1, 0)
		if pointsErr != nil {
			return errors.NewInternalServerError("database transaction add new points error")
		}
	} else {
		// Update User Points
		_, pointsErr := tx.ExecContext(ctx, queryUpdateUserPoints, userPoints+1, 0, userID)
		if pointsErr != nil {
			return errors.NewInternalServerError("database transaction add points error")
		}
	}	

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return errors.NewInternalServerError("database transaction commit error")
	}
	return nil
}

func (receipt *Receipt) Save2() *errors.RestErr {
	stmt, err := receipts_db.Client.Prepare(queryInsertReceipt)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	currentTime := time.Now()
	date_added := currentTime.Format("20060102150405")
	insertResult, saveErr := stmt.Exec(receipt.ModelYear, receipt.Make, receipt.Vin, receipt.FirstName, receipt.LastName, receipt.State, date_added)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	receiptID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	receipt.ID = receiptID
	return nil
}

func (receipt *Receipt) GetAllCount() (int64, *errors.RestErr) {
	var count int64

	row := receipts_db.Client.QueryRow(queryGetAllCount)
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.NewInternalServerError("database error")
	}

	return count, nil
}

func (receipt *Receipt) GetAllCountToday() (int64, *errors.RestErr) {
	var count int64

	row := receipts_db.Client.QueryRow(queryGetAllCountToday)
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.NewInternalServerError("database error")
	}

	return count, nil
}
