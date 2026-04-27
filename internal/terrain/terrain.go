// Package terrain defines the possible terrains, their passability,
// and how to convert them to/from ASCII characters.
package terrain

import "fmt"

type Terrain uint8

const (
	Ant Terrain = iota
	Empty
	Food
	Goal
	Mound
	Path
	Query
	Self
	Wall
	Water
)

func (t Terrain) IsPassable() bool {
	switch t {
	case Ant, Empty, Food, Goal, Mound:
		return true
	case Path, Query, Self, Wall, Water:
		return false
	default:
		panic(fmt.Sprintf("Impossible terrain (isPassable): %d", t))
	}
}

func CharToTerrain(b byte) Terrain {
	switch b {
	case 'a':
		return Ant
	case ' ':
		return Empty
	case 'f':
		return Food
	case 'G':
		return Goal
	case 'O':
		return Mound
	case 'x':
		return Path
	case '.':
		return Query
	case '*':
		return Self
	case 'D':
		return Wall
	case '%':
		return Water
	default:
		panic(fmt.Sprintf("Impossible terrain (CharToTerrain): %d", b))
	}
}

func (t Terrain) ToChar() byte {
	switch t {
	case Ant:
		return 'a'
	case Empty:
		return ' '
	case Food:
		return 'f'
	case Goal:
		return 'G'
	case Mound:
		return 'O'
	case Path:
		return 'x'
	case Query:
		return '.'
	case Self:
		return '*'
	case Wall:
		return 'D'
	case Water:
		return '%'
	default:
		panic(fmt.Sprintf("Impossible terrain (toChar): %d", t))
	}
}
