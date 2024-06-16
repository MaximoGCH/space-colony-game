package db

import (
	"github.com/MaximoGCH/space-colony-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type StructureType int

const (
	HouseLv1 = iota
	HouseLv2
	HouseLv3
	// only for recipes
	AllHouses
	Rock
	Tree
)

type Structure struct {
	Type                StructureType
	Name                string
	Image               *ebiten.Image
	CardDropNumber      int
	CardDropGroupNumber int
}

type StructureDb map[StructureType]*Structure

func createStructureDatabase(assets *assets.Assets) StructureDb {
	return StructureDb{
		HouseLv1: &Structure{
			Type:                HouseLv1,
			Name:                "House Lv1",
			Image:               assets.GetSprite("embed/sprites/house-lv-1"),
			CardDropNumber:      3,
			CardDropGroupNumber: 1,
		},
		HouseLv2: &Structure{
			Type:                HouseLv1,
			Name:                "House Lv2",
			Image:               assets.GetSprite("embed/sprites/house-lv-2"),
			CardDropNumber:      3,
			CardDropGroupNumber: 2,
		},
		HouseLv3: &Structure{
			Type:                HouseLv1,
			Name:                "Human Lv3",
			Image:               assets.GetSprite("embed/sprites/house-lv-3"),
			CardDropNumber:      3,
			CardDropGroupNumber: 3,
		},
		Rock: &Structure{
			Type:                Rock,
			Name:                "Rock",
			Image:               assets.GetSprite("embed/sprites/rock"),
			CardDropNumber:      2,
			CardDropGroupNumber: 1,
		},
		Tree: &Structure{
			Type:                Tree,
			Name:                "Tree",
			Image:               assets.GetSprite("embed/sprites/tree"),
			CardDropNumber:      2,
			CardDropGroupNumber: 1,
		},
	}
}
