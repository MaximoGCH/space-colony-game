package shapes

import "github.com/hajimehoshi/ebiten/v2"

func (point Point) ToImageOptions() *ebiten.DrawImageOptions {
	imageOptions := &ebiten.DrawImageOptions{}
	imageOptions.GeoM.Translate(float64(point.X), float64(point.Y))
	return imageOptions
}

func FromMousePosition() Point {
	x, y := ebiten.CursorPosition()
	return Point{
		X: x,
		Y: y,
	}
}
