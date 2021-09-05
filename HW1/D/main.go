package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	students := 0
	fmt.Scanln(&students)
	reader := bufio.NewReader(os.Stdin)
	inputString, _:= reader.ReadString('\n')

	result := FindSchoolPlace(students, inputString)
	fmt.Print(result)
}

func FindSchoolPlace(size int, input string ) string{
	temp := strings.Trim(input, "\n")
	data := strings.Split(temp, " ")

	var responce int
	length := len(data)
	middle := length / 2
	if (length % 2) == 0{
		a, _ := strconv.Atoi(data[middle])
		b, _ := strconv.Atoi(data[middle - 1])
		responce = (a + b) / 2

	} else {
		responce, _ = strconv.Atoi(data[middle])
	}

	return strconv.Itoa(responce) + "\n"
}
