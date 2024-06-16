package systems

import (
	"fmt"

	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/custom_text"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func UpdateInventory(globalState *state.GlobalState) {
	mousePosition := shapes.FromMousePosition()

	for _, item := range globalState.GameState.Inventory {

		if globalState.GameState.NextDayTransitionPhase != 0 {
			// game is in next day transition
			continue
		}

		intersect := shapes.PointIntersectRectangle(mousePosition, item.Bounds)

		if item.Amount >= 1 && intersect && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			globalState.GameState.Inventory.RemoveResource(item.Type, 1)
			globalState.GameState.ResourceCard = &models.ResourceCard{
				Type:        item.Type,
				Position:    mousePosition,
				MouseOffset: mousePosition.PointSub(item.Bounds.Point),
			}
		}
	}
}

func DrawInventory(globalState *state.GlobalState, screen *ebiten.Image) {
	cardImage := globalState.Assets.GetSprite("embed/sprites/card")
	for _, item := range globalState.GameState.Inventory {
		screen.DrawImage(cardImage, item.Bounds.ToImageOptions())

		if item.Type == db.Empty {
			continue
		}

		resourceImage := globalState.Db.Resources[item.Type].BigImage
		screen.DrawImage(resourceImage, item.Bounds.ToImageOptions())

		custom_text.DrawOutlineText(screen,
			fmt.Sprintf("x%v", item.Amount),
			item.Bounds.PointAdd(shapes.Point{X: 16, Y: 38}),
			globalState.Assets.GetFont("embed/fonts/Kubasta"),
		)
	}
}
