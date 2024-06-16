package systems

import (
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateLostCardSystem(globalState *state.GlobalState) {
	for i := 0; i < len(globalState.GameState.LostResourceCard); i++ {
		card := globalState.GameState.LostResourceCard[i]

		card.VelocityX += float32(card.GoTo.X-card.Position.X) / 100
		card.VelocityY += float32(card.GoTo.Y-card.Position.Y) / 100

		card.VelocityX /= 1.1
		card.VelocityY /= 1.1

		card.Position = card.Position.PointAdd(shapes.Point{
			X: int(card.VelocityX), Y: int(card.VelocityY),
		})

		if shapes.PointIntersectRectangle(card.Position.PointAdd(shapes.Point{
			X: models.CardWidth / 2, Y: models.CardHeight / 5,
		}), card.GoTo) {
			globalState.GameState.LostResourceCard.RemoveIndex(i, globalState.GameState.Inventory)
			i--
		}
	}
}

func DrawLostCardCard(globalState *state.GlobalState, screen *ebiten.Image) {
	cardImage := globalState.Assets.GetSprite("embed/sprites/card")

	for _, card := range globalState.GameState.LostResourceCard {
		resourceImage := globalState.Db.Resources[card.Type].BigImage

		screen.DrawImage(cardImage, card.Position.ToImageOptions())
		screen.DrawImage(resourceImage, card.Position.ToImageOptions())
	}
}
