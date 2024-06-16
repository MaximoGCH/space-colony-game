package expanded_render

import "github.com/hajimehoshi/ebiten/v2"

func NewExpandableSprite(image *ebiten.Image, width int, height int) *ebiten.Image {
	sourceImageSize := image.Bounds().Size()
	newImage := ebiten.NewImage(width, height)

	for i := 0; i < width; i += sourceImageSize.X {
		for j := 0; j < height; j += sourceImageSize.Y {
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(float64(i), float64(j))
			newImage.DrawImage(image, opt)
		}
	}

	return newImage
}
