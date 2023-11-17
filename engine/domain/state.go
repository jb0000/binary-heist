package domain

type Cell struct {
	IsPassable bool
	contains   map[string]string
}
type Level [][]Cell
type State struct {
	Blueprint Level
}

func NewCoord(x, y int) Coord {
	return Coord{X: x, Y: y}
}

type Coord struct {
	X int
	Y int
}
