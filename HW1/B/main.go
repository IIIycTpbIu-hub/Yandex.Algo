package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input,"\n")
	strArr := strings.Split(input, " ")

	data := make([]int, len(strArr))

	for i, str := range strArr{
		data[i], _ = strconv.Atoi(str)
	}

	result := Find(data[0], data[1], data[2])

	fmt.Println(result)
}

func Find(total int, start int, end int) int{
	var min, max int
	if start < end{
		min = start
		max = end
	} else {
		min = end
		max = start
	}
	preResult := max - min
	middle := total / 2

	if preResult <= middle{
		return preResult - 1
	} else {
		return total + min - max - 1
	}
	return 0
}
