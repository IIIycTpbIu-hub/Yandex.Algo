package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\\n")

	strArr := strings.Split(input, " ")
	var intArr = make([]int, len(strArr))

	for i, str := range strArr{
		intArr[i], _ = strconv.Atoi(str)
	}

	result := Process(intArr[0], intArr[1], intArr[2])

	fmt.Println(result)
}

func Process(x int, y int, z int) int{
	if x > 12 || y > 12 || x == y{
		return 1
	} else {
		return 0
	}
}
