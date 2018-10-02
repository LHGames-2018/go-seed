package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	var gamemap Map
	var data string

	data = "[[{1},{3},{1}],[{},{},{4,5000,1}],[{},{},{}]]"
	gamemap.UnmarshalJSON([]byte(data))
	assert.Equal(t, gamemap.Tiles[0][0], Tile{Position: Point{X: 0, Y: 0}, Type: TileWall})
	assert.Equal(t, gamemap.Tiles[0][1], Tile{Position: Point{X: 1, Y: 0}, Type: TileLava})
	assert.Equal(t, gamemap.Tiles[0][2], Tile{Position: Point{X: 2, Y: 0}, Type: TileWall})
	assert.Equal(t, gamemap.Tiles[1][0], Tile{Position: Point{X: 0, Y: 1}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[1][1], Tile{Position: Point{X: 1, Y: 1}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[1][2], Tile{Position: Point{X: 2, Y: 1}, Type: TileResource})
	assert.Equal(t, gamemap.Tiles[2][0], Tile{Position: Point{X: 0, Y: 2}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[2][1], Tile{Position: Point{X: 1, Y: 2}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[2][2], Tile{Position: Point{X: 2, Y: 2}, Type: TileEmpty})

	data = "[[{4,4000,1,1},{3},{1}],[{},{},{4,5000}],[{},{},{}]]"
	gamemap.UnmarshalJSON([]byte(data))
	assert.Equal(t, gamemap.Tiles[0][0], Tile{Position: Point{X: 0, Y: 0}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[0][1], Tile{Position: Point{X: 1, Y: 0}, Type: TileLava})
	assert.Equal(t, gamemap.Tiles[0][2], Tile{Position: Point{X: 2, Y: 0}, Type: TileWall})
	assert.Equal(t, gamemap.Tiles[1][0], Tile{Position: Point{X: 0, Y: 1}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[1][1], Tile{Position: Point{X: 1, Y: 1}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[1][2], Tile{Position: Point{X: 2, Y: 1}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[2][0], Tile{Position: Point{X: 0, Y: 2}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[2][1], Tile{Position: Point{X: 1, Y: 2}, Type: TileEmpty})
	assert.Equal(t, gamemap.Tiles[2][2], Tile{Position: Point{X: 2, Y: 2}, Type: TileEmpty})
}
