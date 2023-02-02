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
