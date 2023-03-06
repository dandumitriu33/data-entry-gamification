package service

import (
	"context"
	"data-entry-gamification/model"
	"data-entry-gamification/utils/errors"

	"log"
)

func CreateReceipt(ctx context.Context, receipt model.Receipt, user model.User) (*model.Receipt, *errors.RestErr) {
	if err := receipt.Validate(); err != nil {
		return nil, err
	}

	var receiptDAO model.ReceiptDAO
	if err := receiptDAO.Save(ctx, user.ID, receipt); err != nil {
		return nil, err
	}

	return &receipt, nil
}

func GetAllCount() (int64, *errors.RestErr) {
	var receipt model.Receipt
	count, err := receipt.GetAllCount()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetAllCountToday() (int64, *errors.RestErr) {
	var receipt model.Receipt
	count, err := receipt.GetAllCountToday()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetUnverifiedReceipt() (model.Receipt, *errors.RestErr) {
	result := &model.ReceiptDAO{}

	resultReceipt, err := result.GetUnverifiedReceipt()
	if err != nil {
		log.Println("Error getting Unverified Receipt from DAO:", err)
		return resultReceipt, err
	}

	return resultReceipt, nil
}

func UpdateReceipt(receipt model.Receipt) (*model.Receipt, *errors.RestErr) {
	if err := receipt.UpdateReceipt(); err != nil {
		return nil, err
	}

	return &receipt, nil
}
