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
	FailedRun     = status.FailedRun
	SuccessfulRun = status.SuccessfulRun
	goal          = terrain.Goal
	path          = terrain.Path
	query         = terrain.Query
	self          = terrain.Self
)

type (
	Breadcrumb       = coord.Breadcrumb
	Coordinate       = coord.Coordinate
	Crumb            = coord.Crumb
	PathingGrid      = interp.PathingGrid
	PathingMapData   = interp.PathingMapData
	PathingMapString = interp.PathingMapString
	RunResult        = status.RunResult
	Source           = coord.Source
)

type PathingMap struct {
	Grid PathingGrid
}

func NewCrumb(to Coordinate, from Breadcrumb) Crumb {
	return Crumb{To: to, From: from}
}

func NewSource(coord Coordinate) Source {
	return Source{Coord: coord}
}

func (pmap PathingMap) maxCoord() Coordinate {
	return slices.MaxFunc(slices.Collect(maps.Keys(pmap.Grid)), coord.Compare)
}

func (pmap PathingMap) Height() uint {
	return pmap.maxCoord().Y + 1
}

func (pmap PathingMap) Width() uint {
	return pmap.maxCoord().X + 1
}

func (pmap PathingMap) String() string {
	if len(pmap.Grid) > 0 {
		coords := slices.Collect(maps.Keys(pmap.Grid))
		slices.SortFunc(coords, coord.Compare)
		maxCoord := coords[len(coords)-1]

		lines := make([]string, 0, maxCoord.Y)
		strBuffer := new(strings.Builder)

		for i, coord := range coords {
			strBuffer.WriteByte(pmap.Grid[coord].ToChar())
			if (i+1)%int(maxCoord.X+1) == 0 { //nolint:gosec // realistic coords are too small to overflow
				lines = append(lines, strBuffer.String())
				strBuffer.Reset()
			}
		}

		linesWithBorders := make([]string, 0, len(lines))
		slices.Reverse(lines) // Recall that we reoriented the coordinates when reading --Jason B. (4/28/26)
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
		pmap.Grid[coord] = path
	}
}

func (pmap PathingMap) MarkAsGoal(coord Coordinate) {
	pmap.Grid[coord] = goal
}

func (pmap PathingMap) Step(prev Coordinate, next Coordinate) {
	pmap.Grid[prev] = query
	pmap.Grid[next] = self
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
		terrain, isOK := pmap.Grid[c]
		if isOK && terrain.IsPassable() {
			out = append(out, c)
		}
	}

	return out
}
