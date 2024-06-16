package expanded_render

import (
	"fmt"
	"image"

	"github.com/MaximoGCH/space-colony-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewNineSliceSprite(assets *assets.Assets, assetDir string, width int, height int) *ebiten.Image {
	newImage := ebiten.NewImage(width, height)

	spriteData := assets.GetSpriteData(assetDir)
	if spriteData.Slices == nil || len(spriteData.Slices.Horizontal) < 2 || len(spriteData.Slices.Vertical) < 2 {
		panic(fmt.Sprintf("Error, can not render slices without a slice objet in the sprite data json %v", assetDir))
	}

	topLeftSprite := assets.GetSpriteSlices(assetDir, 0, 0)
	topCenterSprite := assets.GetSpriteSlices(assetDir, 1, 0)
	topRightSprite := assets.GetSpriteSlices(assetDir, 2, 0)

	middleLeftSprite := assets.GetSpriteSlices(assetDir, 0, 1)
	middleCenterSprite := assets.GetSpriteSlices(assetDir, 1, 1)
	middleRightSprite := assets.GetSpriteSlices(assetDir, 2, 1)

	bottomLeftSprite := assets.GetSpriteSlices(assetDir, 0, 2)
	bottomCenterSprite := assets.GetSpriteSlices(assetDir, 1, 2)
	bottomRightSprite := assets.GetSpriteSlices(assetDir, 2, 2)

	topLeftSize := topLeftSprite.Bounds().Size()
	bottomRightSize := bottomRightSprite.Bounds().Size()
	centerBoundsSize := image.Point{
		X: width - topLeftSize.X - bottomRightSize.X,
		Y: height - topLeftSize.Y - bottomRightSize.Y,
	}

	// top

	opTopLeft := &ebiten.DrawImageOptions{}
	newImage.DrawImage(topLeftSprite, opTopLeft)

	expandableTopCenterSprite := NewExpandableSprite(topCenterSprite, centerBoundsSize.X, topLeftSize.Y)
	opTopCenter := &ebiten.DrawImageOptions{}
	opTopCenter.GeoM.Translate(float64(topLeftSize.X), 0)
	newImage.DrawImage(expandableTopCenterSprite, opTopCenter)

	opTopRight := &ebiten.DrawImageOptions{}
	opTopRight.GeoM.Translate(float64(topLeftSize.X+centerBoundsSize.X), 0)
	newImage.DrawImage(topRightSprite, opTopRight)

	// middle

	expandableMiddleLeftSprite := NewExpandableSprite(middleLeftSprite, topLeftSize.X, centerBoundsSize.Y)
	opMiddleLeft := &ebiten.DrawImageOptions{}
	opMiddleLeft.GeoM.Translate(0, float64(topLeftSize.Y))
	newImage.DrawImage(expandableMiddleLeftSprite, opMiddleLeft)

	expandableMiddleCenterSprite := NewExpandableSprite(middleCenterSprite, centerBoundsSize.X, centerBoundsSize.Y)
	opMiddleCenter := &ebiten.DrawImageOptions{}
	opMiddleCenter.GeoM.Translate(float64(topLeftSize.X), float64(topLeftSize.Y))
	newImage.DrawImage(expandableMiddleCenterSprite, opMiddleCenter)

	expandableMiddleRightSprite := NewExpandableSprite(middleRightSprite, bottomRightSize.X, centerBoundsSize.Y)
	opMiddleRight := &ebiten.DrawImageOptions{}
	opMiddleRight.GeoM.Translate(float64(topLeftSize.X+centerBoundsSize.X), float64(topLeftSize.Y))
	newImage.DrawImage(expandableMiddleRightSprite, opMiddleRight)

	// bottom

	opBottomLeft := &ebiten.DrawImageOptions{}
	opBottomLeft.GeoM.Translate(0, float64(topLeftSize.Y+centerBoundsSize.Y))
	newImage.DrawImage(bottomLeftSprite, opBottomLeft)

	expandableBottomCenterSprite := NewExpandableSprite(bottomCenterSprite, centerBoundsSize.X, bottomRightSize.Y)
	opBottomCenter := &ebiten.DrawImageOptions{}
	opBottomCenter.GeoM.Translate(float64(topLeftSize.X),
		float64(topLeftSize.Y+centerBoundsSize.Y))
	newImage.DrawImage(expandableBottomCenterSprite, opBottomCenter)

	opBottomRight := &ebiten.DrawImageOptions{}
	opBottomRight.GeoM.Translate(float64(topLeftSize.X+centerBoundsSize.X),
		float64(topLeftSize.Y+centerBoundsSize.Y))
	newImage.DrawImage(bottomRightSprite, opBottomRight)

	return newImage
}
