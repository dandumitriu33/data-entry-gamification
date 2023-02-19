package string_utils

import "strings"

func StringContainsOnlyUppercaseCharacters(input string) bool {
	return input == strings.ToUpper(input)
}

func Contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}