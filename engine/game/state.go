package game

type cell struct {
	isPassable bool
	contains map[string]string
}
type blueprint [][]cell
type state struct{
	level blueprint
}