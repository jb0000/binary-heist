package domain

type Cell struct {
	isPassable bool
	contains   map[string]string
}
type Level [][]Cell
type State struct {
	Blueprint Level
}
type Coord struct {
	x int
	y int
}
