package main

import "fmt"

type Point struct {
	x, y int
}

func scale(point Point, scale int) Point {
	return Point{x: point.x * scale, y: point.y * scale}
}

func getX(point Point) int {
	return point.x
}

func getY(point *Point) int {
	return point.y
}

func setX(point Point, x int) {
	point.x = x
}

func setY(point *Point, y int) {
	point.y = y
}

func main() {
	var point = Point{x: 1, y: 2}
	fmt.Println(point)           //	{1 2}
	fmt.Println(scale(point, 2)) //	{2 4}

	fmt.Println(getX(point))  // 1
	fmt.Println(getY(&point)) //	2

	setX(point, 2)
	setY(&point, 4)

	fmt.Println(point) //	{1 4}

	point1 := &Point{1, 2}
	point2 := new(Point)
	*point2 = Point{1, 2}

	fmt.Println(point1) //	&{1 2}
	fmt.Println(point2) //	&{1 2}
}
