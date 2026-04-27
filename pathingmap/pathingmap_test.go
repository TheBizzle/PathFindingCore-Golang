package pathingmap

import (
	"slices"
	"testing"
)

var baseMapStr = PathingMapString{Contents: " DGD | DDD |%%%% |DD %%|*D  %", Delim: "|"}

func gridFromString(s string) PathingMap {
	pms := PathingMapString{Contents: s, Delim: "|"}
	return PathingMap{grid: pms.AsPMD().Grid}
}

func TestNeighborsOf(t *testing.T) {
	base := PathingMap{grid: baseMapStr.AsPMD().Grid}
	testNeighbors("neighborsOf 1", base, Coordinate{X: 9001, Y: 9001}, []Coordinate{}, t)
	testNeighbors("neighborsOf 2", base, Coordinate{X: 2, Y: 0}, []Coordinate{{X: 2, Y: 1}, {X: 3, Y: 0}}, t)
	testNeighbors("neighborsOf 3", base, Coordinate{X: 3, Y: 0}, []Coordinate{{X: 2, Y: 0}}, t)
}

func TestStep(t *testing.T) {
	base := PathingMap{grid: baseMapStr.AsPMD().Grid}
	base.Step(Coordinate{X: 2, Y: 0}, Coordinate{X: 2, Y: 1})
	testMap("step", base, gridFromString(" DGD | DDD |%%%% |DD*%%|*D. %"), t)
}

func TestMarkAsGoal(t *testing.T) {
	base := PathingMap{grid: baseMapStr.AsPMD().Grid}
	base.MarkAsGoal(Coordinate{X: 2, Y: 0})
	testMap("markAsGoal 1", base, gridFromString(" DGD | DDD |%%%% |DD %%|*DG %"), t)

	base2 := PathingMap{grid: baseMapStr.AsPMD().Grid}
	base2.MarkAsGoal(Coordinate{X: 2, Y: 1})
	testMap("markAsGoal 2", base2, gridFromString(" DGD | DDD |%%%% |DDG%%|*D  %"), t)
}

func TestInsertPath(t *testing.T) {
	base1 := PathingMap{grid: baseMapStr.AsPMD().Grid}
	base1.InsertPath([]Coordinate{{X: 4, Y: 3}, {X: 4, Y: 2}})
	testMap("insertPath 1", base1, gridFromString(" DGD | DDDx|%%%%x|DD %%|*D  %"), t)

	base2 := PathingMap{grid: baseMapStr.AsPMD().Grid}
	base2.InsertPath([]Coordinate{{X: 2, Y: 1}, {X: 2, Y: 0}, {X: 3, Y: 0}})
	testMap("insertPath 2", base2, gridFromString(" DGD | DDD |%%%% |DDx%%|*Dxx%"), t)

	base3 := PathingMap{grid: baseMapStr.AsPMD().Grid}
	base3.InsertPath([]Coordinate{})
	testMap("insertPath 3", base3, gridFromString(" DGD | DDD |%%%% |DD %%|*D  %"), t)
}

func testMap(desc string, actual PathingMap, expected PathingMap, t *testing.T) {
	if actual.String() != expected.String() {
		t.Errorf("%s | Expected:\n%s\nActual:\n%s", desc, expected.String(), actual.String())
	}
}

func testNeighbors(desc string, pmap PathingMap, coord Coordinate, expected []Coordinate, t *testing.T) {
	actual := pmap.NeighborsOf(coord)
	if !slices.Equal(actual, expected) {
		t.Errorf("%s | Expected: %v | Actual: %v", desc, expected, actual)
	}
}
