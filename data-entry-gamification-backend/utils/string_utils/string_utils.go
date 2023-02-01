package string_utils

import "strings"

func StringContainsOnlyUppercaseCharacters(input string) bool {
	return input == strings.ToUpper(input)
}