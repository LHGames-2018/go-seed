////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"strconv"
)

import "github.com/prataprc/goparsec"

type Map struct {
	RelativeTo Point
	Tiles      [][]ITile
	Resources  []*Resource
}

func (u *Map) UnmarshalJSON(data []byte) error {
	/* parser's atoms */
	bracketOpen := parsec.Atom(`[`, "BRACKET_OPEN")
	bracketClose := parsec.Atom(`]`, "BRACKET_CLOSE")
	braceOpen := parsec.Atom(`{`, "BRACE_OPEN")
	braceClose := parsec.Atom(`}`, "BRACE_CLOSE")
	comma := parsec.Atom(`,`, "COMMA")

	/* definition of the parser */
	ast := parsec.NewAST("MAP_PARSER", 100)
	tileNumber := parsec.Token("\\d*\\.?\\d+", "TILE_NUMBER")
	tileContent := ast.Many("TILE_CONTENT", nil, tileNumber, comma)
	tileMaybe := ast.Maybe("TILE_MAYBE", nil, tileContent)
	tileItem := ast.And("TILE_ITEM", nil, braceOpen, tileMaybe, braceClose)
	rowContent := ast.Many("ROW_CONTENT", nil, tileItem, comma)
	rowItem := ast.And("ROW_ITEM", nil, bracketOpen, rowContent, bracketClose)
	mapContent := ast.Many("MAP_CONTENT", nil, rowItem, comma)
	mapItem := ast.And("MAP_ITEM", nil, bracketOpen, mapContent, bracketClose)

	/* parse the input string */
	scanner := parsec.NewScanner(data)
	root, err := ast.Parsewith(mapItem, scanner)
	if err == nil {
		return errors.New("Failed to deserialize map.")
	}

	/* parse the resulting tree */
	resources := []*Resource{}
	rows := [][]ITile{}
	for i, rowNode := range root.GetChildren()[1].GetChildren() {
		row := []ITile{}
		for j, tileNode := range rowNode.GetChildren()[1].GetChildren() {
			/* get the nodes*/
			tile := tileNode.GetChildren()[1]
			nodes := tile.GetChildren()

			/* create position relative to this map */
			position := Point{X: i, Y: j}

			/* nothing means it's empty */
			if len(nodes) == 0 {
				tile := Tile{Type: TileEmpty, Position: position}
				row = append(row, &tile)
				continue
			}

			/* try to decode the first node */
			codef, err := strconv.ParseFloat(nodes[0].GetValue(), 32)
			if err != nil {
				fmt.Println("Failed to parse `" + nodes[0].GetValue() + "`.")
				tile := Tile{Type: TileEmpty, Position: position}
				row = append(row, &tile)
				continue
			}
			code := int(codef)

			/* create a resource if it is one */
			if code == TileResource {
				/* make sure there is 3 fields */
				if len(nodes) != 3 {
					fmt.Println("Wrong number of field to a resource.")
					tile := Tile{Type: TileEmpty, Position: position}
					row = append(row, &tile)
					continue
				}

				/* parse the second field */
				maxf, err := strconv.ParseFloat(nodes[1].GetValue(), 32)
				if err != nil {
					fmt.Println("Failed to parse `" + nodes[1].GetValue() + "`.")
					tile := Tile{Type: TileEmpty, Position: position}
					row = append(row, &tile)
					continue
				}
				max := int(maxf)

				/* parse the third field */
				density, err := strconv.ParseFloat(nodes[2].GetValue(), 32)
				if err != nil {
					fmt.Println("Failed to parse `" + nodes[2].GetValue() + "`.")
					tile := Tile{Type: TileEmpty, Position: position}
					row = append(row, &tile)
					continue
				}

				/* add the resource */
				resource := Resource{
					Position:  position,
					Type:      TileResource,
					Remaining: max,
					Density:   float32(density),
				}
				resources = append(resources, &resource)
				row = append(row, &resource)
				continue
			}

			/* the code should be the same as the enumeration in tile.go */
			ttile := Tile{Type: code, Position: position}
			row = append(row, &ttile)
		}
		rows = append(rows, row)
	}

	/* update the struct's fields */
	u.Tiles = rows
	u.Resources = resources
	u.RelativeTo = Point{X: 0, Y: 0}

	return nil
}

func (m *Map) Print() {
	/* mapping tiles to bytes */
	mapping := map[int]byte{
		TileEmpty:    ' ',
		TileWall:     'X',
		TileHouse:    'H',
		TileLava:     'L',
		TileResource: 'R',
		TileShop:     'S',
		TilePlayer:   'P',
	}

	/* print all tile */
	for _, row := range m.Tiles {
		for _, tile := range row {
			fmt.Print(string(mapping[tile.GetType()]) + "  ")
		}
		fmt.Println("")
	}
}

func (m *Map) SetRelativeTo(x int, y int) {
	for i, row := range m.Tiles {
		for j, tile := range row {
			newX := x + tile.GetPosition().X - m.RelativeTo.X
			newY := y + tile.GetPosition().Y - m.RelativeTo.Y
			m.Tiles[i][j].SetPosition(Point{X: newX, Y: newY})
		}
	}
	m.RelativeTo = Point{X: x, Y: y}
}

func (m *Map) GetTile(x int, y int) ITile {
	/* get position relative to the array */
	x = x - m.RelativeTo.X
	y = y - m.RelativeTo.Y

	/* make sure the x coordinate is inside the map */
	if x < 0 || x >= len(m.Tiles) {
		return nil
	}

	/* get the column */
	column := m.Tiles[x]

	/* make sure the y coordinate is inside the map */
	if y < 0 || y >= len(column) {
		return nil
	}

	return column[y]
}
