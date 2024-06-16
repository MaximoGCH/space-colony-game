package systems

import (
	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/expanded_render"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func UpdateCardDrop(globalState *state.GlobalState, cardDrop *models.CardDrop) {
	if globalState.GameState.NextDayTransitionPhase != 0 {
		// game is in next day transition
		return
	}

	mousePosition := shapes.FromMousePosition()

	// set resources
	intersect := shapes.PointIntersectRectangle(mousePosition, cardDrop.Bounds)

	// set resource from card
	if intersect && inpututil.IsMouseButtonJustReleased(
		ebiten.MouseButton0) && globalState.GameState.ResourceCard != nil {

		if cardDrop.Type != db.Empty {
			globalState.GameState.LostResourceCard.Add(
				&models.ResourceCard{
					Type:        cardDrop.Type,
					Position:    cardDrop.Bounds.Point,
					MouseOffset: shapes.Point{},
				},
				globalState.GameState.Inventory,
				true,
			)
		}

		cardDrop.Type = globalState.GameState.ResourceCard.Type
		globalState.GameState.ResourceCard = nil
	}

	// remove resource and add card
	if intersect && cardDrop.Type != db.Empty && inpututil.IsMouseButtonJustPressed(
		ebiten.MouseButton0) {
		globalState.GameState.ResourceCard = &models.ResourceCard{
			Type:        cardDrop.Type,
			Position:    mousePosition,
			MouseOffset: mousePosition.PointSub(cardDrop.Bounds.Point),
		}
		cardDrop.Type = db.Empty
	}
}

func DrawCardDrop(globalState *state.GlobalState, screen *ebiten.Image, cardDrop *models.CardDrop) {
	// draw card drop
	cardDropImage := expanded_render.NewNineSliceSprite(
		globalState.Assets, "embed/sprites/dashed-line",
		models.CardWidth,
		models.CardHeight)

	screen.DrawImage(cardDropImage, cardDrop.Bounds.ToImageOptions())

	// draw card drop resource image
	if cardDrop.Type == db.Empty {
		return
	}

	cardDropResourceImage := globalState.Db.Resources[cardDrop.Type].BigImage

	screen.DrawImage(cardDropResourceImage, cardDrop.Bounds.ToImageOptions())
}
