package service

import (
	"data-entry-gamification/model"
	"data-entry-gamification/utils/errors"
)

func CreateReceipt(receipt model.Receipt) (*model.Receipt, *errors.RestErr) {
	if err := receipt.Validate(); err != nil {
		return nil, err
	}

	if err := receipt.Save(); err != nil {
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
