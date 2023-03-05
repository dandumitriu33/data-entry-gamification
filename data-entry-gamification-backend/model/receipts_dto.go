package model

import (
	"data-entry-gamification/utils/errors"
	"data-entry-gamification/utils/string_utils"

	"strings"
	"time"
)

func (receipt *Receipt) Validate() *errors.RestErr {
	receipt.FirstName = strings.TrimSpace(receipt.FirstName)
	if receipt.FirstName == "" {
		return errors.NewBadRequestError("invalid receipt first name")
	}
	receipt.LastName = strings.TrimSpace(receipt.LastName)
	if receipt.LastName == "" {
		return errors.NewBadRequestError("invalid receipt last name")
	}
	receipt.Make = strings.TrimSpace(receipt.Make)
	if receipt.Make == "" {
		return errors.NewBadRequestError("invalid receipt make")
	}
	currentYear := time.Now().Year()
	if receipt.ModelYear > currentYear || receipt.ModelYear < 1800 {
		return errors.NewBadRequestError("invalid receipt model year")
	}
	receipt.State = strings.TrimSpace(receipt.State)
	if len(receipt.State) != 2 || !string_utils.StringContainsOnlyUppercaseCharacters(receipt.State) {
		return errors.NewBadRequestError("invalid receipt state")
	}
	receipt.Vin = strings.TrimSpace(receipt.Vin)
	if len(receipt.Vin) != 17 || !string_utils.StringContainsOnlyUppercaseCharacters(receipt.Vin) {
		return errors.NewBadRequestError("invalid receipt vin")
	}
	return nil
}
