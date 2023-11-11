package loader

import "github.com/jb0000/binary-heist/engine/domain"

type DefaultLoader struct {
}

func (d DefaultLoader) Load() domain.State {
	layout := `
		 XXXXXXX
		 X     X
		 X  XXXX
		XX  X
		 X  X
		 X  X
		 XXXX`
	bp := levelFromString(layout)

	return domain.State{
		Blueprint: bp,
	}
}

func levelFromString(layout string) domain.Level {
	return initLevel(7, 7)
}

func initLevel(x, y int) domain.Level {
	bp := make(domain.Level, x-1)
	for i := x - 1; i >= 0; i-- {
		bp[i] = make([]domain.Cell, y-1)
	}

	return bp
}
