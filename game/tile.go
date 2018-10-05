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

type ITile interface {
	GetPosition() Point
	GetType() int
	SetPosition(Point)
}

type Tile struct {
	Position Point
	Type     int
}

func (t *Tile) GetPosition() Point {
	return t.Position
}

func (t *Tile) GetType() int {
	return t.Type
}

func (t *Tile) SetPosition(point Point) {
	t.Position = point
}

type Resource struct {
	Position  Point
	Type      int
	Remaining int
	Density   float32
}

func (r *Resource) GetPosition() Point {
	return r.Position
}

func (r *Resource) GetType() int {
	return r.Type
}

func (r *Resource) SetPosition(point Point) {
	r.Position = point
}
