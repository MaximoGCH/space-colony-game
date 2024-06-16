package models

import (
	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
)

type ResourceCard struct {
	Type        db.ResourceType
	Position    shapes.Point
	MouseOffset shapes.Point
}
