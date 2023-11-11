package domain

type Cell struct {
	IsPassable bool
	contains   map[string]string
}
type Level [][]Cell
type State struct {
	Blueprint Level
}
type Coord struct {
	X int
	Y int
}
