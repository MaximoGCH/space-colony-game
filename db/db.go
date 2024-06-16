package db

import "github.com/MaximoGCH/space-colony-game/assets"

type Db struct {
	Resources       ResourceDb
	StructureList   StructureList
	Structures      StructureDb
	ResourceList    ResourceList
	Recipes         RecipeDb
	UnsortedRecipes RecipeDb
}

func CreateDb(assets *assets.Assets) *Db {
	sortedRecipes, unsortedRecipes := CreateRecipeDatabase()

	return &Db{
		Resources:       CreateResourceDatabase(assets),
		StructureList:   CreateStructureList(),
		Structures:      CreateStructureDatabase(assets),
		ResourceList:    CreateResourceList(),
		Recipes:         sortedRecipes,
		UnsortedRecipes: unsortedRecipes,
	}
}
