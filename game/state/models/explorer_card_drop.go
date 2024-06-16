package models

import (
	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
)

const (
	ExplorerCardDropLen    = 6
	ExplorerCardDropXStart = 850
	ExplorerCardDropYStart = 96
)

type ExplorerCardDrop [ExplorerCardDropLen]*CardDrop

func CreateExplorerCardDrop() ExplorerCardDrop {
	cardDrop := ExplorerCardDrop{}

	for i := 0; i < ExplorerCardDropLen; i++ {
		cardDrop[i] = &CardDrop{
			Type: db.Empty,
			Bounds: shapes.Rectangle{
				Point: shapes.Point{
					X: ExplorerCardDropXStart,
					Y: ExplorerCardDropYStart + i*CardDropGroupSize,
				},
				Size: shapes.Size{
					Width:  CardWidth,
					Height: CardHeight,
				},
			},
		}
	}

	return cardDrop
}
