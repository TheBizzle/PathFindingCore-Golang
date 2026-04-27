package interpreter

import (
	"slices"
	"testing"

	"github.com/TheBizzle/PathFindingCore-Golang/internal/coordinate"
	"github.com/TheBizzle/PathFindingCore-Golang/internal/terrain"
)

var (
	Empty = terrain.Empty
	Mound = terrain.Mound
	Wall  = terrain.Wall
	Water = terrain.Water
)

var (
	m5x5p1  = "DDDDD|DG  D|D   D|D  *D|DDDDD"
	m5x5p2  = " DGD | DDD |%%%% |DD %%|*D  %"
	trblmkr = " %  *|OG% %|%    |"
)

func TestInterpreter(t *testing.T) {
	testInterpreter(t, "Simple grid", "*G", Coordinate{X: 0, Y: 0}, Coordinate{X: 1, Y: 0}, Coordinate{X: 1, Y: 0}, []Terrain{Self, Goal})
	testInterpreter(t, "One-line grid 1", "*      G", Coordinate{X: 0, Y: 0}, Coordinate{X: 7, Y: 0}, Coordinate{X: 7, Y: 0}, []Terrain{Self, Empty, Empty, Empty, Empty, Empty, Empty, Goal})
	testInterpreter(t, "One-line grid 2", "G      *", Coordinate{X: 7, Y: 0}, Coordinate{X: 0, Y: 0}, Coordinate{X: 7, Y: 0}, []Terrain{Goal, Empty, Empty, Empty, Empty, Empty, Empty, Self})
	testInterpreter(t, "One-line grid 3", "G %D%  *", Coordinate{X: 7, Y: 0}, Coordinate{X: 0, Y: 0}, Coordinate{X: 7, Y: 0}, []Terrain{Goal, Empty, Water, Wall, Water, Empty, Empty, Self})
	testInterpreter(t, "Simple vertical grid", "*|G", Coordinate{X: 0, Y: 1}, Coordinate{X: 0, Y: 0}, Coordinate{X: 0, Y: 1}, []Terrain{Goal, Self})
	testInterpreter(t, "One-line vert grid 1", "*| | | |G", Coordinate{X: 0, Y: 4}, Coordinate{X: 0, Y: 0}, Coordinate{X: 0, Y: 4}, []Terrain{Goal, Empty, Empty, Empty, Self})
	testInterpreter(t, "One-line vert grid 2", "G| | | |*", Coordinate{X: 0, Y: 0}, Coordinate{X: 0, Y: 4}, Coordinate{X: 0, Y: 4}, []Terrain{Self, Empty, Empty, Empty, Goal})
	testInterpreter(t, "One-line vert grid 3", "G| |%|D|*", Coordinate{X: 0, Y: 0}, Coordinate{X: 0, Y: 4}, Coordinate{X: 0, Y: 4}, []Terrain{Self, Wall, Water, Empty, Goal})
	testInterpreter(t, "2x2 grid 1", "G | *", Coordinate{X: 1, Y: 0}, Coordinate{X: 0, Y: 1}, Coordinate{X: 1, Y: 1}, []Terrain{Empty, Goal, Self, Empty})
	testInterpreter(t, "2x2 grid 2", "G*|  ", Coordinate{X: 1, Y: 1}, Coordinate{X: 0, Y: 1}, Coordinate{X: 1, Y: 1}, []Terrain{Empty, Goal, Empty, Self})
	testInterpreter(t, "2x2 grid 3", "G*|DD", Coordinate{X: 1, Y: 1}, Coordinate{X: 0, Y: 1}, Coordinate{X: 1, Y: 1}, []Terrain{Wall, Goal, Wall, Self})
	testInterpreter(t, "2x2 grid 4", "DD|*G", Coordinate{X: 0, Y: 0}, Coordinate{X: 1, Y: 0}, Coordinate{X: 1, Y: 1}, []Terrain{Self, Wall, Goal, Wall})
	testInterpreter(t, "5x5 grid 1", m5x5p1, Coordinate{X: 3, Y: 1}, Coordinate{X: 1, Y: 3}, Coordinate{X: 4, Y: 4}, []Terrain{Wall, Wall, Wall, Wall, Wall, Wall, Empty, Empty, Goal, Wall, Wall, Empty, Empty, Empty, Wall, Wall, Self, Empty, Empty, Wall, Wall, Wall, Wall, Wall, Wall})
	testInterpreter(t, "5x5 grid 2", m5x5p2, Coordinate{X: 0, Y: 0}, Coordinate{X: 2, Y: 4}, Coordinate{X: 4, Y: 4}, []Terrain{Self, Wall, Water, Empty, Empty, Wall, Wall, Water, Wall, Wall, Empty, Empty, Water, Wall, Goal, Empty, Water, Water, Wall, Wall, Water, Water, Empty, Empty, Empty})
	testInterpreter(t, "Past troublemaker", trblmkr, Coordinate{X: 4, Y: 2}, Coordinate{X: 1, Y: 1}, Coordinate{X: 4, Y: 2}, []Terrain{Water, Mound, Empty, Empty, Goal, Water, Empty, Water, Empty, Empty, Empty, Empty, Empty, Water, Self})
}

func testInterpreter(t *testing.T, desc string, strGrid string, start Coordinate, goal Coordinate, maxCoord Coordinate, terrains []Terrain) {
	t.Helper()
	pms := PathingMapString{Contents: strGrid, Delim: "|"}
	actual := pms.AsPMD()

	if actual.Start != start {
		t.Errorf("%s | Start: expected %v, got %v", desc, start, actual.Start)
	}
	if actual.Goal != goal {
		t.Errorf("%s | Goal: expected %v, got %v", desc, goal, actual.Goal)
	}

	expectedGrid := gridFromList(maxCoord, terrains)
	actualGrid := gridToSortedSlice(actual.Grid)
	if !slices.Equal(actualGrid, expectedGrid) {
		t.Errorf("%s | Grid mismatch:\n  expected: %v\n  actual:   %v", desc, expectedGrid, actualGrid)
	}
}

type coordTerrain struct {
	coord   Coordinate
	terrain Terrain
}

func gridToSortedSlice(grid PathingGrid) []coordTerrain {
	pairs := make([]coordTerrain, 0, len(grid))
	for coord, t := range grid {
		pairs = append(pairs, coordTerrain{coord, t})
	}
	slices.SortFunc(pairs, func(a, b coordTerrain) int {
		return coordinate.Compare(a.coord, b.coord)
	})
	return pairs
}

func gridFromList(maxCoord Coordinate, terrains []Terrain) []coordTerrain {
	pairs := make([]coordTerrain, 0, len(terrains))
	i := 0
	for x := uint(0); x <= maxCoord.X; x++ {
		for y := uint(0); y <= maxCoord.Y; y++ {
			pairs = append(pairs, coordTerrain{Coordinate{X: x, Y: y}, terrains[i]})
			i++
		}
	}
	slices.SortFunc(pairs, func(a, b coordTerrain) int {
		return coordinate.Compare(a.coord, b.coord)
	})
	return pairs
}
