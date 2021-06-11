package main

import "testing"

func Test_BracketMatcher_DataDrivenTests(t *testing.T) {
	testCases := []struct {
		testName       string
		testString     string
		expectedResult bool
	}{
		{
			"Fail: One Open", "{", false,
		},
		{
			"Fail: One Closed", "}", false,
		},
		{
			"Fail: Closed before Open", "}{", false,
		},
		{
			"Success: Matching", "{}", true,
		},
		{
			"Success: No brackets", "ad", true,
		},
		{
			"Success: Matching with extras", "goodstring { anything works!@K }!kd,", true,
		},
	}

	for _, testCase := range testCases {
		if MatchBrackets(testCase.testString) != testCase.expectedResult {
			t.Errorf("TestCase %v failed on %v", testCase.testName, testCase.testString)
		}
	}
}
