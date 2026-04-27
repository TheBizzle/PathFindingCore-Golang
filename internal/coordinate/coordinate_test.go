package coordinate

import (
	"slices"
	"testing"
)

func TestBreadcrumbs(t *testing.T) {
	testIt("Simple source 1", Source{Coordinate{X: 0, Y: 0}}, []Coordinate{{X: 0, Y: 0}}, t)
	testIt("Simple source 2", Source{Coordinate{X: 1, Y: 1}}, []Coordinate{{X: 1, Y: 1}}, t)
	testIt("Simple source 3", Source{Coordinate{X: 3, Y: 8}}, []Coordinate{{X: 3, Y: 8}}, t)
	testIt("Two-item crumb", Crumb{Coordinate{X: 0, Y: 0}, Source{Coordinate{X: 3, Y: 8}}}, []Coordinate{{X: 3, Y: 8}, {X: 0, Y: 0}}, t)
	testIt("Three-item crumb", Crumb{Coordinate{X: 1, Y: 7}, Crumb{Coordinate{X: 0, Y: 0}, Source{Coordinate{X: 3, Y: 8}}}}, []Coordinate{{X: 3, Y: 8}, {X: 0, Y: 0}, {X: 1, Y: 7}}, t)
}

func testIt(desc string, crumbs Breadcrumb, coords []Coordinate, t *testing.T) {
	actual := crumbs.Array()
	expected := coords
	if !slices.Equal(actual, expected) {
		t.Errorf("%s | Expected: %v | Actual: %v", desc, expected, actual)
	}
}
