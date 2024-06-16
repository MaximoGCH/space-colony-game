package systems

import (
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func UpdateResourceCard(globalState *state.GlobalState) {
	if globalState.GameState.ResourceCard == nil {
		return
	}

	globalState.GameState.ResourceCard.Position =
		shapes.FromMousePosition().PointSub(globalState.GameState.ResourceCard.MouseOffset)

	if inpututil.IsMouseButtonJustReleased(
		ebiten.MouseButton0) {
		globalState.GameState.LostResourceCard.Add(
			globalState.GameState.ResourceCard,
			globalState.GameState.Inventory,
			false,
		)
		globalState.GameState.ResourceCard = nil
	}
}

func DrawResourceCard(globalState *state.GlobalState, screen *ebiten.Image) {
	if globalState.GameState.ResourceCard == nil {
		return
	}

	cardImage := globalState.Assets.GetSprite("embed/sprites/card")
	resourceImage := globalState.Db.Resources[globalState.GameState.ResourceCard.Type].BigImage

	screen.DrawImage(cardImage, globalState.GameState.ResourceCard.Position.ToImageOptions())
	screen.DrawImage(resourceImage, globalState.GameState.ResourceCard.Position.ToImageOptions())
}
