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
	{"1 2 2003", "0"},
	{"2 29 2008", "1"},
	{"12 12 2008", "1"},
	{"29 2 2008", "1"},
}

func Test(t *testing.T){
	var input []string
	for _, testCase := range testCases{
		input = strings.Split(testCase.input, " ")
		var inputInts = make([]int, len(input))

		for i, str := range input{
			inputInts[i], _ = strconv.Atoi(str)
		}

		var result = strconv.Itoa(Process(inputInts[0], inputInts[1], inputInts[2]))

		if result != testCase.output{
			t.Error("For input ", testCase.input, "expected ", testCase.output, "but was ", result)
		}
	}
}