package db

import "github.com/MaximoGCH/space-colony-game/assets"

type Db struct {
	Resources    ResourceDb
	Structures   StructureDb
	ResourceList ResourceList
	Recipes      RecipeDb
}

func CreateDb(assets *assets.Assets) *Db {
	return &Db{
		Resources:    CreateResourceDatabase(assets),
		Structures:   createStructureDatabase(assets),
		ResourceList: CreateResourceList(),
		Recipes:      CreateRecipeDatabase(),
	}
}
