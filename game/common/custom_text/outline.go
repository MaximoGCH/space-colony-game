package custom_text

import (
	"image/color"

	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func DrawOutlineText(image *ebiten.Image, txt string, position shapes.Point, font *font.Face) {
	bounds := text.BoundString(*font, txt)

	centeredPos := position.PointAdd(shapes.Point{
		X: -bounds.Min.X - bounds.Dx()/2,
		Y: -bounds.Min.Y - bounds.Dy()/2,
	})

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			if i == 0 && j == 0 {
				continue
			}

			text.Draw(image,
				txt,
				*font,
				centeredPos.X+i,
				centeredPos.Y+j,
				color.Black,
			)
		}
	}

	text.Draw(image,
		txt,
		*font,
		centeredPos.X,
		centeredPos.Y,
		color.White,
	)
}
