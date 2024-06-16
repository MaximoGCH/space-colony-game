package systems

import (
	"github.com/MaximoGCH/space-colony-game/game/common/animations/easing"
	"github.com/MaximoGCH/space-colony-game/game/common/custom_text"
	"github.com/MaximoGCH/space-colony-game/game/common/expanded_render"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	NotificationTime = 500
)

func UpdateNotifications(globalState *state.GlobalState) {
	for i := 0; i < len(globalState.GameState.Notifications); i++ {
		notification := globalState.GameState.Notifications[i]
		if !notification.Remove {
			notification.Height.Update(true, 10, easing.InOutCubic)
		} else {
			notification.Height.Update(false, 10, easing.InOutCubic)
		}

		if notification.Remove && notification.Height.Value == 0 {
			globalState.GameState.Notifications.RemoveIndex(i)
			i--
		}

	}

	globalState.GameState.NotificationsJustCleared = false
	if globalState.GameState.ClearNotifications {
		globalState.GameState.Notifications.Add(models.Text,
			"Click to continue")
		globalState.GameState.ClearNotificationsCheck = true
		globalState.GameState.ClearNotifications = false
	}

	if globalState.GameState.ClearNotificationsCheck {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			globalState.GameState.NotificationsJustCleared = true
			globalState.GameState.ClearNotificationsCheck = false
			globalState.GameState.Notifications.Clear()
		}
	}
}

func DrawNotifications(globalState *state.GlobalState, screen *ebiten.Image) {
	font := globalState.Assets.GetFont("embed/fonts/Kubasta")
	screenCenter := globalState.ScreenSize.ToPoint().ConstDiv(2)

	yOffset := 0
	for _, notification := range globalState.GameState.Notifications {
		textBounds := text.BoundString(*font, notification.Text)
		size := shapes.Size{
			Width:  textBounds.Dx() + 32,
			Height: textBounds.Dy() + 20,
		}

		bounds := shapes.Rectangle{
			Size: shapes.Size{
				Width:  size.Width,
				Height: (size.Height * notification.Height.Value) / 100,
			},
			Point: shapes.Point{
				X: screenCenter.X - (size.Width / 2),
				Y: yOffset + 20,
			},
		}

		yOffset += bounds.Height

		if bounds.Height <= 0 {
			continue
		}

		image := expanded_render.NewNineSliceSprite(globalState.Assets, "embed/sprites/notification",
			bounds.Width, bounds.Height)

		custom_text.DrawOutlineText(image,
			notification.Text,
			bounds.Size.Center(),
			font,
		)

		screen.DrawImage(image, bounds.ToImageOptions())
	}
}
