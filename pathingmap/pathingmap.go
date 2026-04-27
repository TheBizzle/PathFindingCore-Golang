// Package pathingmap defines structures and functions for mutating and printing `PathingMap` objects
package pathingmap

import (
	"maps"
	"slices"
	"strings"

	coord "github.com/TheBizzle/PathFindingCore-Golang/internal/coordinate"
	interp "github.com/TheBizzle/PathFindingCore-Golang/internal/interpreter"
	status "github.com/TheBizzle/PathFindingCore-Golang/internal/status"
	terrain "github.com/TheBizzle/PathFindingCore-Golang/internal/terrain"
)

const (
	goal  = terrain.Goal
	path  = terrain.Path
	query = terrain.Query
	self  = terrain.Self
)

type (
	Breadcrumb       = coord.Breadcrumb
	Coordinate       = coord.Coordinate
	PathingGrid      = interp.PathingGrid
	PathingMapData   = interp.PathingMapData
	PathingMapString = interp.PathingMapString
	RunResult        = status.RunResult
)

type PathingMap struct {
	grid PathingGrid
}

func (pmap PathingMap) maxCoord() Coordinate {
	return slices.MaxFunc(slices.Collect(maps.Keys(pmap.grid)), coord.Compare)
}

func (pmap PathingMap) Height() uint {
	return pmap.maxCoord().Y + 1
}

func (pmap PathingMap) Width() uint {
	return pmap.maxCoord().X + 1
}

func (pmap PathingMap) String() string {
	if len(pmap.grid) > 0 {
		coords := slices.Collect(maps.Keys(pmap.grid))
		slices.SortFunc(coords, coordinate.Compare)
		maxCoord := coords[len(coords)-1]

		lines := make([]string, 0, maxCoord.Y)
		strBuffer := new(strings.Builder)

		for i, coord := range coords {
			strBuffer.WriteByte(pmap.grid[coord].ToChar())
			if (i+1)%int(maxCoord.X+1) == 0 { //nolint:gosec // realistic coords are too small to overflow
				lines = append(lines, strBuffer.String())
				strBuffer.Reset()
			}
		}

		linesWithBorders := make([]string, 0, len(lines))
		for _, line := range lines {
			linesWithBorders = append(linesWithBorders, "|"+line+"|\n")
		}

		border := strings.Repeat("-", int(maxCoord.X)+1) //nolint:gosec // realistic coords are too small to overflow
		topBorder := "+" + border + "+\n"
		bottomBorder := "+" + border + "+"

		return topBorder + strings.Join(linesWithBorders, "") + bottomBorder

	} else {
		return ""
	}
}

func (pmap PathingMap) InsertPath(coords []Coordinate) {
	for _, coord := range coords {
		pmap.grid[coord] = path
	}
}

func (pmap PathingMap) MarkAsGoal(coord Coordinate) {
	pmap.grid[coord] = goal
}

func (pmap PathingMap) Step(prev Coordinate, next Coordinate) {
	pmap.grid[prev] = query
	pmap.grid[next] = self
}

func (pmap PathingMap) NeighborsOf(coord Coordinate) []Coordinate {
	options := []Coordinate{
		{X: coord.X, Y: coord.Y + 1},
		{X: coord.X, Y: coord.Y - 1}, // Integers will underflow, but it's okay --Jason B. (4/26/26)
		{X: coord.X + 1, Y: coord.Y},
		{X: coord.X - 1, Y: coord.Y}, // Underflow here, too
	}

	out := make([]Coordinate, 0, len(options))

	for _, c := range options {
		terrain, isOK := pmap.grid[c]
		if isOK && terrain.IsPassable() {
			out = append(out, c)
		}
	}

	return out
}
