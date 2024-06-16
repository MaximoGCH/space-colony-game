package systems

import (
	"github.com/MaximoGCH/space-colony-game/game/common/custom_text"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateExplorerCardDrop(globalState *state.GlobalState) {
	for _, cardDrop := range globalState.GameState.ExplorerCardDrop {
		UpdateCardDrop(globalState, cardDrop)
	}
}

func DrawExplorerCardDrop(globalState *state.GlobalState, screen *ebiten.Image) {
	custom_text.DrawOutlineText(screen,
		"Exploration",
		shapes.Point{X: models.ExplorerCardDropXStart - 16, Y: models.ExplorerCardDropYStart - 12},
		globalState.Assets.GetFont("embed/fonts/Kubasta"),
	)

	for _, cardDrop := range globalState.GameState.ExplorerCardDrop {
		DrawCardDrop(globalState, screen, cardDrop)
	}
}
