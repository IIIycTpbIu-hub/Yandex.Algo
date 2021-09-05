package main

import (
	"strconv"
	"testing"
)

type testCase struct {
	students int
	buildings string
	output string
}

var testCases = []testCase{

	{4, "1 2 3 4\n" , "2\n"},
	{3, "-1 0 1\n", "0\n"},
	{5, "-10 -5 0 5 10\n", "0\n"},
	{5, "-5 -4 -3 -2 -1\n", "-3\n"},
	{5, "-9 -1 8 9 10\n", "8\n"},
}

func Test0(t *testing.T)  {
	for _, testCase := range testCases{
		result := FindSchoolPlace(testCase.students, testCase.buildings)

		if result != testCase.output{
			t.Error(
				strconv.Itoa(testCase.students) + "\n",
				testCase.buildings,
				"expected", testCase.output,
				"actual", result)

		}
	}
}
