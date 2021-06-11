package main

import "strings"

func MatchBrackets(stringToTest string) bool {
	bracketCount := 0

	for _, char := range strings.Split(stringToTest, "") {
		if char == "{" {
			bracketCount++
		} else if char == "}" {
			bracketCount--
		}

		if bracketCount < 0 {
			return false
		}
	}

	return bracketCount == 0
}
