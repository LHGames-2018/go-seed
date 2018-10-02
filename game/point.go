////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

type Point struct {
	X int
	Y int
}

type JSONPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (json JSONPoint) Point() Point {
	return Point{
		X: json.X,
		Y: json.Y,
	}
}
