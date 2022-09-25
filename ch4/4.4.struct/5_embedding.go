package main

import "fmt"

//type Circle struct {
//	X, Y, Radius int
//}

//type Wheel struct {
//	X, Y, Radius, Spokes int
//}

type Point2 struct {
	X, Y int
}

type Circle2 struct {
	Center Point2
	Radius int
}

type Wheel2 struct {
	Circle Circle2
	Spokes int
}

type Point3 struct {
	X, Y int
}

type Circle3 struct {
	Point3
	Radius int
}

type Wheel3 struct {
	Circle3
	Spokes int
}

func main() {
	var w Wheel2
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	var wh Wheel3
	wh.X = 8
	wh.Y = 8
	wh.Radius = 5
	wh.Spokes = 20

	//wh2 := Wheel3{8, 8, 5, 20} // compile error: unknown fields
	//wh2 := Wheel3{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	w2 := Wheel3{Circle3{Point3{8, 8}, 5}, 20}
	w2 = Wheel3{
		Circle3: Circle3{
			Point3: Point3{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w2) //	Wheel3{Circle3:Circle3{Point3:Point3{X:8, Y:8}, Radius:5}, Spokes:20}
	w2.X = 42
	fmt.Printf("%#v\n", w2) // Wheel{Circle3:Circle3{Point3:Point3{X:42, Y:8}, Radius:5}, Spokes:20}
}
