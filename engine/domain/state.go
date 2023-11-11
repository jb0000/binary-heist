package domain

type Cell struct {
	isPassable bool
	contains map[string]string
}
type Blueprint [][]cell
type State struct{
	level blueprint
}
type Coord struct{
	x int
	y int
}