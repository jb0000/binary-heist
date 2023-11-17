package application

import "github.com/jb0000/binary-heist/engine/domain"

type Loader interface {
	Load() domain.State
}
type game struct {
	loader Loader
}

func NewGame(loader Loader) game {
	return game{
		loader: loader,
	}
}
func (g game) StartGame() {
	_ = g.loader.Load()
}
