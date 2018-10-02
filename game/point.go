////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type JSONPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Point) ToJSON() string {
	return fmt.Sprintf("{X:%d,Y:%d}", p.X, p.Y)
}

func (json JSONPoint) Point() Point {
	return Point{
		X: json.X,
		Y: json.Y,
	}
}
