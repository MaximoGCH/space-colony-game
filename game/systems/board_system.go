package systems

import (
	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/animations/easing"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateBoard(globalState *state.GlobalState) {
	for structureI := 0; structureI < len(globalState.GameState.Board); structureI++ {
		for structureJ := 0; structureJ < len(globalState.GameState.Board[structureI]); structureJ++ {
			// draw structure
			structure := globalState.GameState.Board[structureI][structureJ]

			if structure == nil {
				continue
			}

			// start animation
			structure.AnimationOffset.Update(
				true, 5, easing.OutSine,
			)

			// card drops
			for _, cardDropGroup := range structure.CardDrop {
				for _, cardDrop := range cardDropGroup {
					UpdateCardDrop(globalState, cardDrop)
				}
			}
		}
	}
}

func DrawBoard(globalState *state.GlobalState, screen *ebiten.Image) {
	// draw board line
	// cardDropImage := expanded_render.NewNineSliceSprite(
	// 	globalState.Assets, "embed/sprites/dashed-line",
	// 	(models.BoardSizeW-1)*models.GridSpace,
	// 	(models.BoardSizeH)*models.GridSpace+30,
	// )

	// screen.DrawImage(cardDropImage, shapes.Point{
	// 	X: models.BoardStartX + models.GridSize + 20,
	// 	Y: models.BoardStartY - 15,
	// }.ToImageOptions())

	// draw structures
	for i := 0; i < len(globalState.GameState.Board); i++ {
		for j := 0; j < len(globalState.GameState.Board[i]); j++ {
			// draw structure
			structure := globalState.GameState.Board[i][j]

			if structure == nil {
				continue
			}

			var easingOffset shapes.Point
			if structure.Type == db.HouseLv1 {
				easingOffset = shapes.Point{
					X: 0,
					Y: structure.AnimationOffset.MaxValue - structure.AnimationOffset.EasingValue,
				}
			} else {
				easingOffset = shapes.Point{
					X: -structure.AnimationOffset.MaxValue + structure.AnimationOffset.EasingValue,
					Y: 0,
				}
			}

			structureInfo := globalState.Db.Structures[structure.Type]
			screen.DrawImage(structureInfo.Image, structure.Position.PointSub(easingOffset).ToImageOptions())

			if structure.AnimationOffset.MaxValue != structure.AnimationOffset.EasingValue {
				continue
			}

			for _, cardDropGroup := range structure.CardDrop {
				for _, cardDrop := range cardDropGroup {
					DrawCardDrop(globalState, screen, cardDrop)
				}
			}
		}
	}
}
