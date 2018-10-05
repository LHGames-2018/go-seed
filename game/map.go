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
	Tiles      [][]Tile
	Resources  []Resource
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
	resources := []Resource{}
	rows := [][]Tile{}
	for j, rowNode := range root.GetChildren()[1].GetChildren() {
		row := []Tile{}
		for i, tileNode := range rowNode.GetChildren()[1].GetChildren() {
			/* get the nodes*/
			tile := tileNode.GetChildren()[1]
			nodes := tile.GetChildren()

			/* create position relative to this map */
			position := Point{X: i, Y: j}

			/* nothing means it's empty */
			if len(nodes) == 0 {
				row = append(row, Tile{Type: TileEmpty, Position: position})
				continue
			}

			/* try to decode the first node */
			codef, err := strconv.ParseFloat(nodes[0].GetValue(), 32)
			if err != nil {
				fmt.Println("Failed to parse `" + nodes[0].GetValue() + "`.")
				row = append(row, Tile{Type: TileEmpty, Position: position})
				continue
			}
			code := int(codef)

			/* create a resource if it is one */
			if code == TileResource {
				/* make sure there is 3 fields */
				if len(nodes) != 3 {
					fmt.Println("Wrong number of field to a resource.")
					row = append(row, Tile{Type: TileEmpty, Position: position})
					continue
				}

				/* parse the second field */
				maxf, err := strconv.ParseFloat(nodes[1].GetValue(), 32)
				if err != nil {
					fmt.Println("Failed to parse `" + nodes[1].GetValue() + "`.")
					row = append(row, Tile{Type: TileEmpty, Position: position})
					continue
				}
				max := int(maxf)

				/* parse the third field */
				density, err := strconv.ParseFloat(nodes[2].GetValue(), 32)
				if err != nil {
					fmt.Println("Failed to parse `" + nodes[2].GetValue() + "`.")
					row = append(row, Tile{Type: TileEmpty, Position: position})
					continue
				}

				/* add the resource */
				resource := Resource{Position: position, Remaining: max, Density: float32(density)}
				resources = append(resources, resource)
			}

			/* the code should be the same as the enumeration in tile.go */
			row = append(row, Tile{Type: code, Position: position})
		}
		rows = append(rows, row)
	}

	/* update the struct's fields */
	u.Tiles = rows
	u.Resources = resources

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
			fmt.Print(string(mapping[tile.Type]) + " ")
		}
		fmt.Println("")
	}
}

func (m *Map) SetRelativeTo(x int, y int) {
	for j, row := range m.Tiles {
		for i, tile := range row {
			tile.Position.X = i + x
			tile.Position.Y = j + y
		}
	}
	m.RelativeTo.X = x
	m.RelativeTo.Y = y
}
