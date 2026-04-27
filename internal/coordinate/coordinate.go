// Package coordinate defines the data structures for coordinates and breadcrumbs.
package coordinate

import "cmp"

type Coordinate struct {
	X, Y uint
}

func Compare(a, b Coordinate) int {
	if comparison := cmp.Compare(a.Y, b.Y); comparison != 0 {
		return comparison
	} else {
		return cmp.Compare(a.X, b.X)
	}
}

type Breadcrumb interface {
	isBreadcrumb()
	Array() []Coordinate // Me no likey using arrays for this --Jason. B (4/25/26)
}

type Crumb struct {
	To   Coordinate
	From Breadcrumb
}

type Source struct {
	Coord Coordinate
}

var (
	_ Breadcrumb = Source{}
	_ Breadcrumb = Crumb{}
)

func (s Source) isBreadcrumb() {}
func (c Crumb) isBreadcrumb()  {}

func (s Source) Array() []Coordinate { return []Coordinate{s.Coord} }
func (c Crumb) Array() []Coordinate {
	return append(c.From.Array(), c.To) // Dislike. --Jason B. (4/25/26)
}
