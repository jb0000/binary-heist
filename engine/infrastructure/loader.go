package infrastructur

import (
	"strings"

	"github.com/jb0000/binary-heist/engine/domain"
)

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
	rows := strings.Split(layout, "\n")
	var maxX int
	for _, row := range rows {
		if len(row) > maxX {
			maxX = len(row)
		}
	}

	bp := initLevel(maxX, len(rows))

	coords = getPassableCoordsFromRows(rows)

	bp = makeRoomsPassable(bp, coords)
	
	return bp
}

func getPassableCoordsFromRows(rows []string) []domain.Coord {
	var coords []domain.Coord
	for y, row := range rows {
		for x, cell := range row {
			if cell != ' ' {
				coords = append(coords, domain.Coord{x, y})
			}
		}
	}

	return coords
}

func makeRoomsPassable(bp domain.Blueprint, coords []domain.Coord)domain.Blueprint{
 for _, coord := range choords{
   bp[coord.X][coord.Y].IsPassable = true
 }
 return bp
}

func initLevel(x, y int) domain.Level {
	bp := make(domain.Level, x)
	for i := x - 1; i >= 0; i-- {
		bp[i] = make([]domain.Cell, y)
	}

	return bp
}
