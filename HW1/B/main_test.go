package main

import (
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	input string
	output string
}

var testCases = []testCase{
	{"100 5 6", "0"},
	{"10 1 9", "1"},
	{"10 3 9", "3"},
	{"10 3 8", "4"},
	{"7 1 6", "1"},
	{"7 2 5", "2"},
	{"7 1 7", "0"},
	{"100 6 5", "0"},
	{"10 9 1", "1"},
	{"10 9 3", "3"},
	{"10 8 3", "4"},
	{"7 6 1", "1"},
	{"7 5 2", "2"},
	{"7 7 1", "0"},
}

func Test(t *testing.T) {

	for _, testCase := range testCases{
		input := strings.Split(testCase.input, " ")
		numSlice := make([]int, len(input))
		for i, item := range input{
			numSlice[i], _ = strconv.Atoi(item)
		}

		result := Find(numSlice[0], numSlice[1], numSlice[2])

		var expected, _ = strconv.Atoi(testCase.output)
		if result != expected{
			t.Error("For input: ", testCase.input,", expected: ", expected, ",but was: ", result)
		}
	}
}
