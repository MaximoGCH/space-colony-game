package systems

import (
	"github.com/MaximoGCH/space-colony-game/game/common/custom_text"
	"github.com/MaximoGCH/space-colony-game/game/common/expanded_render"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func UpdateButton(button *models.Button) {
	mousePos := shapes.FromMousePosition()
	button.IsJustActive = false
	button.Pressed = false

	if shapes.PointIntersectRectangle(mousePos, button.Bounds) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			button.Pressed = true
		}

		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			button.IsJustActive = true
		}
	}
}

func DrawButton(globalState *state.GlobalState, screen *ebiten.Image, button *models.Button) {
	var imageName string
	var textYOffset int
	if button.Pressed {
		imageName = "embed/sprites/button-2"
		textYOffset = -2
	} else {
		imageName = "embed/sprites/button-1"
		textYOffset = 1
	}

	image := expanded_render.NewNineSliceSprite(globalState.Assets, imageName,
		button.Bounds.Width, button.Bounds.Height)

	screen.DrawImage(image, button.Bounds.ToImageOptions())
	custom_text.DrawOutlineText(screen,
		button.Text,
		button.Bounds.Center().PointAdd(shapes.Point{X: 0, Y: -textYOffset}),
		globalState.Assets.GetFont("embed/fonts/Kubasta"),
	)
}
