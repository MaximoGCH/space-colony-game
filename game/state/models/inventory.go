package models

import (
	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
)

type InventoryItem struct {
	Type   db.ResourceType
	Amount int
	Bounds shapes.Rectangle
}

type Inventory map[db.ResourceType]*InventoryItem

const (
	CardWidth       = 32
	CardHeight      = 48
	CardSpaceWidth  = CardWidth + 4
	CardSpaceHeight = CardHeight + 20
)

func CreateInventory(resources db.ResourceDb, resourceOrder db.ResourceList, screenSize shapes.Size) Inventory {
	// create inventory
	inventory := make(Inventory)
	i := 0
	centerW := screenSize.ToPoint().PointDiv(shapes.Point{X: 2, Y: 1})
	middleW := (len(resources) / 2) * CardSpaceWidth
	for _, resourceType := range resourceOrder {
		inventory[resourceType] = &InventoryItem{
			Type:   resourceType,
			Amount: 0,
			Bounds: shapes.Rectangle{
				Point: centerW.PointAdd(shapes.Point{
					X: -middleW + i*CardSpaceWidth - CardSpaceWidth/2, Y: -CardSpaceHeight,
				}),
				Size: shapes.Size{Width: CardWidth, Height: CardHeight},
			},
		}
		i++
	}

	// start humans and food
	// todo: modify this to allow different planets
	inventory[db.Human].Amount = 3
	inventory[db.Food].Amount = 9

	return inventory
}

func (inventory Inventory) AddResource(resource db.ResourceType, amount int) {
	inventory[resource].Amount += amount
}

func (inventory Inventory) RemoveResource(resource db.ResourceType, amount int) {
	inventory[resource].Amount -= amount
}
