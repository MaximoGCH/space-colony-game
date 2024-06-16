package models

import (
	"math/rand"

	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
)

type LostResourceCard struct {
	ResourceCard
	GoTo      shapes.Rectangle
	Subtract  bool
	VelocityX float32
	VelocityY float32
}

type LostResourceCardList []*LostResourceCard

func (list *LostResourceCardList) Add(resourceCard *ResourceCard, inventory Inventory, animate bool) {
	item := inventory[resourceCard.Type]

	var velX, velY float32 = 0, 0

	if animate {
		velX = rand.Float32()*30 - 15
		velY = rand.Float32()*30 - 15
	}

	*list = append(*list, &LostResourceCard{
		ResourceCard: *resourceCard,
		GoTo:         item.Bounds,
		VelocityX:    velX,
		VelocityY:    velY,
	})
}

func (list *LostResourceCardList) RemoveIndex(index int, inventory Inventory) {
	def := *list
	item := def[index]
	if item.Subtract {
		inventory.RemoveResource(item.Type, 1)
	} else {
		inventory.AddResource(item.Type, 1)
	}
	*list = append(def[:index], def[index+1:]...)
}
