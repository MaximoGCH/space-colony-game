package models

import "github.com/MaximoGCH/space-colony-game/game/common/shapes"

type Button struct {
	Bounds       shapes.Rectangle
	Text         string
	Pressed      bool
	IsJustActive bool
}

func CreateButton(bounds shapes.Rectangle, text string) *Button {
	return &Button{
		Bounds:       bounds,
		Text:         text,
		Pressed:      false,
		IsJustActive: false,
	}
}
