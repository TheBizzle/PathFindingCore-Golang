// Package direction defines the possible directions of movement
package direction

type Direction uint8

//go:generate stringer -type=Direction
const (
	North Direction = iota
	East
	South
	West
)

func Directions() []Direction {
	return []Direction{North, East, South, West}
}
