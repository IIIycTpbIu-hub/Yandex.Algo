package main

import "testing"

type testCase struct {
	d int
	x int
	y int
	result int
}

var testCases = []testCase{
	{5, 1, 1, 0},
	{3, -1, -1, 1},
	{4, 4, 4, 2},
	{4, 2, 2, 0},
}

func Test(t *testing.T){
	for _, testCase := range testCases{
		p := point{testCase.x, testCase.y}
		result := CheckPoint(p, testCase.d)

		if result != testCase.result{
			t.Error("For d:", testCase.d, ", x: ", testCase.x, ", y: ", testCase.y,
				"expected", testCase.result, "but was", result)
		}
	}
}
