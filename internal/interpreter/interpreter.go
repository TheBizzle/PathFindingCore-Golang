// Package interpreter defines structures and functions for reading pathfinding maps from strings
package interpreter

import (
	"fmt"
	"strings"

	"github.com/TheBizzle/PathFindingCore-Golang/internal/coordinate"
	"github.com/TheBizzle/PathFindingCore-Golang/internal/terrain"
)

type (
	Coordinate = coordinate.Coordinate
	Terrain    = terrain.Terrain
)

const (
	Goal = terrain.Goal
	Self = terrain.Self
)

var CharToTerrain = terrain.CharToTerrain

type PathingGrid = map[Coordinate]Terrain

type PathingMapString struct {
	Contents string
	Delim    string
}

type PathingMapData struct {
	Start Coordinate
	Goal  Coordinate
	Grid  PathingGrid
}

func (pms PathingMapString) AsPMD() PathingMapData {
	if pms.Contents != "" {
		strList := strings.Split(strings.TrimSuffix(pms.Contents, pms.Delim), pms.Delim)
		grid := strListToGrid(strList)
		start, goal := findStartAndGoal(grid)
		return PathingMapData{start, goal, grid}
	} else {
		panic("Cannot build map from empty string")
	}
}

func findStartAndGoal(grid PathingGrid) (Coordinate, Coordinate) {
	var start *Coordinate = nil
	var goal *Coordinate = nil

	for k, v := range grid {
		switch v {
		case Self:
			start = &k
		case Goal:
			goal = &k
		}
	}

	if start == nil {
		panic(fmt.Sprintf("No start in given grid: %v", grid))
	} else if goal == nil {
		panic(fmt.Sprintf("No goal in given grid: %v", grid))
	} else {
		return *start, *goal
	}
}

// We start with rows of strings (strs[a][b] => a: 0 = top row, b: 0 = leftmost character) and we need to
// transform it such that it follows normal Cartesian coordinate rules (strs[a][b] => a: 0 = leftmost
// character, b: 0 = bottom row) --Jason B. (4/25/26)
func strListToGrid(strList []string) PathingGrid {
	outMap := PathingGrid{}
	rotated := rotateClockwise(strList)
	for x, col := range rotated {
		for y, val := range col {
			outMap[Coordinate{X: uint(x), Y: uint(y)}] = CharToTerrain(byte(val)) //nolint:gosec // range index is never negative
		}
	}
	return outMap
}

func rotateClockwise(rows []string) []string {
	if len(rows) != 0 {
		numRows := len(rows)
		numCols := len(rows[0])

		output := make([]string, numCols)

		for col := range numCols {
			chars := make([]byte, numRows)
			for row := range numRows {
				chars[row] = rows[numRows-1-row][col]
			}
			output[col] = string(chars)
		}

		return output
	} else {
		return []string{}
	}
}
