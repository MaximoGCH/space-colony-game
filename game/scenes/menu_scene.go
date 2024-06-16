package scenes

import (
	"image/color"

	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/MaximoGCH/space-colony-game/game/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

func StartMenuScene(globalState *state.GlobalState) {
	screenMiddlePos := globalState.ScreenSize.Center()

	globalState.Scene = state.MenuScene
	globalState.GameState = &state.GameState{
		StartButton: models.CreateButton(
			shapes.Rectangle{
				Point: shapes.Point{
					X: screenMiddlePos.X - 160/2,
					Y: screenMiddlePos.Y + 100,
				},
				Size: shapes.Size{
					Width:  160,
					Height: 48,
				},
			},
			"Start colony",
		),
	}
}

func UpdateMenuScene(globalState *state.GlobalState) {
	systems.UpdateButton(globalState.GameState.StartButton)

	if globalState.GameState.StartButton.IsJustActive {
		StartGameScene(globalState)
	}
}

func DrawMenuScene(globalState *state.GlobalState, screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 109,
		G: 141,
		B: 138,
		A: 255,
	})
	screenMiddlePos := globalState.ScreenSize.Center()

	titleImage := globalState.Assets.GetSprite("embed/sprites/tittle")
	screen.DrawImage(titleImage, shapes.Point{
		X: screenMiddlePos.X - 160,
		Y: screenMiddlePos.Y - 180,
	}.ToImageOptions())

	planetImage := globalState.Assets.GetSprite("embed/sprites/planet")
	screen.DrawImage(planetImage, shapes.Point{
		X: screenMiddlePos.X - 64,
		Y: screenMiddlePos.Y - 64,
	}.ToImageOptions())

	systems.DrawButton(globalState, screen, globalState.GameState.StartButton)
}
