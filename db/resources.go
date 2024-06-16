package db

import (
	"github.com/MaximoGCH/space-colony-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type ResourceType int

const (
	Empty = iota
	Human
	Food
	Wood
	Stone
	Axe
	Peak
	Casserole
)

type Resource struct {
	Type       ResourceType
	Name       string
	BigImage   *ebiten.Image
	SmallImage *ebiten.Image
}

type (
	ResourceDb   map[ResourceType]*Resource
	ResourceList []ResourceType
)

// This is the order of the cards in the interface
func CreateResourceList() ResourceList {
	return ResourceList{
		Human,
		Food,
		Wood,
		Stone,
		Axe,
		Peak,
		Casserole,
	}
}

func CreateResourceDatabase(assets *assets.Assets) ResourceDb {
	return ResourceDb{
		Human: &Resource{
			Type:       Human,
			Name:       "Human",
			BigImage:   assets.GetSprite("embed/sprites/human-big"),
			SmallImage: assets.GetSprite("embed/sprites/human-small"),
		},
		Food: &Resource{
			Type:       Food,
			Name:       "Food",
			BigImage:   assets.GetSprite("embed/sprites/food-big"),
			SmallImage: assets.GetSprite("embed/sprites/food-small"),
		},
		Wood: &Resource{
			Type:       Wood,
			Name:       "Wood",
			BigImage:   assets.GetSprite("embed/sprites/wood-big"),
			SmallImage: assets.GetSprite("embed/sprites/wood-small"),
		},
		Stone: &Resource{
			Type:       Stone,
			Name:       "Stone",
			BigImage:   assets.GetSprite("embed/sprites/stone-big"),
			SmallImage: assets.GetSprite("embed/sprites/stone-small"),
		},
		Axe: &Resource{
			Type:       Axe,
			Name:       "Axe",
			BigImage:   assets.GetSprite("embed/sprites/axe-big"),
			SmallImage: assets.GetSprite("embed/sprites/axe-small"),
		},
		Peak: &Resource{
			Type:       Peak,
			Name:       "Peak",
			BigImage:   assets.GetSprite("embed/sprites/peak-big"),
			SmallImage: assets.GetSprite("embed/sprites/peak-small"),
		},
		Casserole: &Resource{
			Type:       Casserole,
			Name:       "Casserole",
			BigImage:   assets.GetSprite("embed/sprites/casserole-big"),
			SmallImage: assets.GetSprite("embed/sprites/casserole-small"),
		},
	}
}
