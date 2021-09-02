package main

import (
	"fmt"
)

func main()  {
	var exitCode, interactor, checker int

	fmt.Scanln(&exitCode)
	fmt.Scanln(&interactor)
	fmt.Scanln(&checker)

	result := ProcessAlgo(exitCode, interactor, checker)

	fmt.Println(result)
}

func ProcessAlgo(exitCode int, interactor int, checker int) int{
	var result = 0
	switch interactor {
	case 0:
		if exitCode!= 0{
			result = 3
		} else {
			result = checker
		}
	case 1:
		result = checker
	case 4:
		if exitCode!= 0{
			result = 3
		} else {
			result = 4
		}
	case 6:
		result = 0
	case 7:
		result = 1

	default:
		result = interactor
	}

	return result
}
