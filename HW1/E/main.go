package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func (a point) distanceTo(b point) float64{
	return math.Sqrt(
			math.Pow(float64(b.x - a.x), 2) +
			math.Pow(float64(b.y - a.y), 2))
}

func main(){
	var d int
	fmt.Scan(&d)

	var pointX = point{}
	fmt.Scanln(&pointX.x, &pointX.y)

	var result = CheckPoint(pointX, d)

	fmt.Println(result)
}

func CheckPoint(pointX point, d int) int{
	var pointA, pointB, pointC point = point{0,0}, point{d, 0}, point{0, d}

	var a int = (pointA.x - pointX.x) * (pointB.y - pointA.y) - (pointB.x - pointA.x) * (pointA.y - pointX.y)
	var b int = (pointB.x - pointX.x) * (pointC.y - pointB.y) - (pointC.x - pointB.x) * (pointB.y - pointX.y)
	var c int = (pointC.x - pointX.x) * (pointA.y - pointC.y) - (pointA.x - pointC.x) * (pointC.y - pointX.y)
	
	if (a >= 0 && b >= 0 && c >= 0) || (a <= 0 && b <= 0 && c <= 0)	{
		return 0
	} else {
		distanceToA := pointA.distanceTo(pointX)
		distanceToB := pointB.distanceTo(pointX)
		distanceToC := pointC.distanceTo(pointX)

		var points = []float64 {distanceToA, distanceToB, distanceToC}

		minVal := distanceToA
		minInd := 0
		for i := 1; i < len(points); i++{
			if points[i] < minVal{
				minVal = points[i]
				minInd = i
			}
		}
		return minInd + 1
	}
}


