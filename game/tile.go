////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

const (
	TileEmpty    = 0
	TileWall     = 1
	TileHouse    = 2
	TileLava     = 3
	TileResource = 4
	TileShop     = 5
	TilePlayer   = 6
)

type Tile struct {
	Position Point
	Type     int
}

type Resource struct {
	Position  Point
	Remaining int
	Density   float32
}
