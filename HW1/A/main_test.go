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
	{"0 0 0 ", "0"},
	{"-1 0 1", "3"},
	{"42 1 6", "6"},
	{"44 7 4", "1"},
	{"1 4 0", "3"},
	{"-3 2 4 ", "2"},
}

func Test(t *testing.T) {
	for _, testCase := range testCases {
		var strings = strings.Split(testCase.input," ")
		ints := make([]int, len(strings))

		for i, s := range strings{
			ints[i], _ = strconv.Atoi(s)
		}

		var result = ProcessAlgo(ints[0], ints[1], ints[2])

		if strconv.Itoa(result) != testCase.output{
			t.Error("Expected ", testCase.output, "but was ", result)
		}
	}
}
