package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	var gamemap Map
	var data string

	/*
	 * The test data for the equivalent map:
	 *
	 *      [1] [ ] [ ]
	 *      [3] [ ] [ ]
	 *      [1] [4] [ ]
	 *
	 */
	data = "[[{1},{3},{1}],[{},{},{4,5000,1.1}],[{},{},{}]]"
	gamemap.UnmarshalJSON([]byte(data))
	assert.Equal(t, gamemap.GetTile(0, 0).GetPosition(), Point{X: 0, Y: 0})
	assert.Equal(t, gamemap.GetTile(1, 0).GetPosition(), Point{X: 1, Y: 0})
	assert.Equal(t, gamemap.GetTile(2, 0).GetPosition(), Point{X: 2, Y: 0})
	assert.Equal(t, gamemap.GetTile(0, 1).GetPosition(), Point{X: 0, Y: 1})
	assert.Equal(t, gamemap.GetTile(1, 1).GetPosition(), Point{X: 1, Y: 1})
	assert.Equal(t, gamemap.GetTile(2, 1).GetPosition(), Point{X: 2, Y: 1})
	assert.Equal(t, gamemap.GetTile(0, 2).GetPosition(), Point{X: 0, Y: 2})
	assert.Equal(t, gamemap.GetTile(1, 2).GetPosition(), Point{X: 1, Y: 2})
	assert.Equal(t, gamemap.GetTile(2, 2).GetPosition(), Point{X: 2, Y: 2})

	assert.Equal(t, gamemap.GetTile(0, 0).GetType(), TileWall)
	assert.Equal(t, gamemap.GetTile(1, 0).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(2, 0).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(0, 1).GetType(), TileLava)
	assert.Equal(t, gamemap.GetTile(1, 1).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(2, 1).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(0, 2).GetType(), TileWall)
	assert.Equal(t, gamemap.GetTile(1, 2).GetType(), TileResource)
	assert.Equal(t, gamemap.GetTile(2, 2).GetType(), TileEmpty)

	assert.Equal(t, len(gamemap.Resources), 1)
	assert.Equal(t, gamemap.Resources[0].Position, Point{X: 1, Y: 2})
	assert.Equal(t, gamemap.Resources[0].Remaining, 5000)
	assert.Equal(t, gamemap.Resources[0].Density, float32(1.1))

	/* change map's relative point*/
	gamemap.SetRelativeTo(20, 40)
	assert.Equal(t, gamemap.GetTile(20, 40).GetPosition(), Point{X: 20, Y: 40})
	assert.Equal(t, gamemap.GetTile(21, 40).GetPosition(), Point{X: 21, Y: 40})
	assert.Equal(t, gamemap.GetTile(22, 40).GetPosition(), Point{X: 22, Y: 40})
	assert.Equal(t, gamemap.GetTile(20, 41).GetPosition(), Point{X: 20, Y: 41})
	assert.Equal(t, gamemap.GetTile(21, 41).GetPosition(), Point{X: 21, Y: 41})
	assert.Equal(t, gamemap.GetTile(22, 41).GetPosition(), Point{X: 22, Y: 41})
	assert.Equal(t, gamemap.GetTile(20, 42).GetPosition(), Point{X: 20, Y: 42})
	assert.Equal(t, gamemap.GetTile(21, 42).GetPosition(), Point{X: 21, Y: 42})
	assert.Equal(t, gamemap.GetTile(22, 42).GetPosition(), Point{X: 22, Y: 42})

	assert.Equal(t, gamemap.GetTile(20, 40).GetType(), TileWall)
	assert.Equal(t, gamemap.GetTile(21, 40).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(22, 40).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(20, 41).GetType(), TileLava)
	assert.Equal(t, gamemap.GetTile(21, 41).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(22, 41).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(20, 42).GetType(), TileWall)
	assert.Equal(t, gamemap.GetTile(21, 42).GetType(), TileResource)
	assert.Equal(t, gamemap.GetTile(22, 42).GetType(), TileEmpty)

	assert.Nil(t, gamemap.GetTile(19, 40))
	assert.Nil(t, gamemap.GetTile(23, 40))
	assert.Nil(t, gamemap.GetTile(20, 39))
	assert.Nil(t, gamemap.GetTile(20, 43))
	assert.Equal(t, len(gamemap.Resources), 1)
	assert.Equal(t, gamemap.Resources[0].Position, Point{X: 21, Y: 42})
	assert.Equal(t, gamemap.Resources[0].Remaining, 5000)
	assert.Equal(t, gamemap.Resources[0].Density, float32(1.1))

	/*
	 * The test data for the equivalent map (with malformed resources):
	 *
	 *      [ ] [ ] [ ]
	 *      [3] [ ] [ ]
	 *      [1] [ ] [ ]
	 *
	 */
	data = "[[{4,4000,1,1},{3},{1}],[{},{},{4,5000}],[{},{},{}]]"
	gamemap.UnmarshalJSON([]byte(data))
	assert.Equal(t, gamemap.GetTile(0, 0).GetPosition(), Point{X: 0, Y: 0})
	assert.Equal(t, gamemap.GetTile(1, 0).GetPosition(), Point{X: 1, Y: 0})
	assert.Equal(t, gamemap.GetTile(2, 0).GetPosition(), Point{X: 2, Y: 0})
	assert.Equal(t, gamemap.GetTile(0, 1).GetPosition(), Point{X: 0, Y: 1})
	assert.Equal(t, gamemap.GetTile(1, 1).GetPosition(), Point{X: 1, Y: 1})
	assert.Equal(t, gamemap.GetTile(2, 1).GetPosition(), Point{X: 2, Y: 1})
	assert.Equal(t, gamemap.GetTile(0, 2).GetPosition(), Point{X: 0, Y: 2})
	assert.Equal(t, gamemap.GetTile(1, 2).GetPosition(), Point{X: 1, Y: 2})
	assert.Equal(t, gamemap.GetTile(2, 2).GetPosition(), Point{X: 2, Y: 2})

	assert.Equal(t, gamemap.GetTile(0, 0).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(1, 0).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(2, 0).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(0, 1).GetType(), TileLava)
	assert.Equal(t, gamemap.GetTile(1, 1).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(2, 1).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(0, 2).GetType(), TileWall)
	assert.Equal(t, gamemap.GetTile(1, 2).GetType(), TileEmpty)
	assert.Equal(t, gamemap.GetTile(2, 2).GetType(), TileEmpty)

	assert.Equal(t, len(gamemap.Resources), 0)
}
