package scenes

import (
	"fmt"
	"image/color"

	"github.com/MaximoGCH/space-colony-game/game/common/custom_text"
	"github.com/MaximoGCH/space-colony-game/game/common/expanded_render"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/MaximoGCH/space-colony-game/game/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

func StartLostScene(globalState *state.GlobalState, days int) {
	screenMiddlePos := globalState.ScreenSize.Center()

	globalState.Scene = state.LostScene
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
			"Restart colony",
		),
		Days: days,
	}
}

func UpdateLostScene(globalState *state.GlobalState) {
	systems.UpdateButton(globalState.GameState.StartButton)

	if globalState.GameState.StartButton.IsJustActive {
		StartGameScene(globalState)
	}
}

func DrawLostScene(globalState *state.GlobalState, screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 109,
		G: 141,
		B: 138,
		A: 255,
	})
	screenMiddlePos := globalState.ScreenSize.Center()

	skullImage := globalState.Assets.GetSprite("embed/sprites/skull")
	screen.DrawImage(skullImage, shapes.Point{
		X: screenMiddlePos.X - 16,
		Y: screenMiddlePos.Y - 120,
	}.ToImageOptions())

	textBack := expanded_render.NewNineSliceSprite(globalState.Assets, "embed/sprites/notification",
		320, 100)

	screen.DrawImage(textBack, shapes.Point{
		X: screenMiddlePos.X - 160,
		Y: screenMiddlePos.Y - 50,
	}.ToImageOptions())

	custom_text.DrawOutlineText(screen,
		fmt.Sprintf("All the inhabitants of your colony have died.\nYou have managed to last %v days.\nI think you can do better.", globalState.GameState.Days),
		shapes.Point{
			X: screenMiddlePos.X,
			Y: screenMiddlePos.Y,
		},
		globalState.Assets.GetFont("embed/fonts/Kubasta"),
	)

	systems.DrawButton(globalState, screen, globalState.GameState.StartButton)
}
