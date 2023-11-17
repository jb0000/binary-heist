package infrastructure

import (
	"strings"

	"github.com/jb0000/binary-heist/engine/domain"
)

type configLoader struct {
	config Config
}

/* `
 XXXXXXX
 X     X
 X  XXXX
XX  X
 X  X
 X  X
 XXXX` */

type Config struct {
	layout string
}

func NewConfigLoader(cfg Config) configLoader {
	return configLoader{
		config: cfg,
	}
}

func (d configLoader) Load() domain.State {
	bp := levelFromString(d.config.layout)

	return domain.State{
		Blueprint: bp,
	}
}

func levelFromString(layout string) domain.Level {
	rows := strings.Split(layout, "\n")
	var maxX int
	for _, row := range rows {
		if len(row) > maxX {
			maxX = len(row)
		}
	}

	bp := initLevel(maxX, len(rows))

	coords := getPassableCoordsFromRows(rows)

	bp = makeRoomsPassable(bp, coords)

	return bp
}

func getPassableCoordsFromRows(rows []string) []domain.Coord {
	var coords []domain.Coord
	for x, row := range rows {
		for y, cell := range row {
			if cell == ' ' {
				coords = append(coords, domain.NewCoord(x, y))
			}
		}
	}

	return coords
}

func makeRoomsPassable(bp domain.Level, coords []domain.Coord) domain.Level {
	for _, coord := range coords {
		bp[coord.X][coord.Y].IsPassable = true
	}
	return bp
}

func initLevel(x, y int) domain.Level {
	bp := make(domain.Level, y)
	for i := y - 1; i >= 0; i-- {
		bp[i] = make([]domain.Cell, x)
	}

	return bp
}
