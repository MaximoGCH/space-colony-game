package shapes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (p *Point) DebugDraw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(p.X), float32(p.Y), 10, color.Black, false)
	vector.DrawFilledCircle(screen, float32(p.X), float32(p.Y), 8, color.White, false)
}

func (p *Rectangle) DebugDraw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(p.X), float32(p.Y), float32(p.Width), float32(p.Height), color.Black, false)
	vector.DrawFilledRect(screen, float32(p.X)+1, float32(p.Y)+1, float32(p.Width)-2,
		float32(p.Height)-2, color.White, false)
}
