package db

import (
	"sort"
)

type Recipe struct {
	NoConsume []ResourceType
	Consume   []ResourceType
	Result    ResourceType
	Amount    int
}

type (
	RecipeDb map[StructureType][]*Recipe
)

func CreateRecipeDatabase() (RecipeDb, RecipeDb) {
	unsortedRecipes := RecipeDb{
		AllHouses: {
			{
				NoConsume: []ResourceType{
					Human,
					Human,
				},
				Consume: []ResourceType{},
				Result:  Human,
				Amount:  1,
			},
			{
				NoConsume: []ResourceType{
					Human,
					Human,
					Human,
					Human,
				},
				Consume: []ResourceType{},
				Result:  Human,
				Amount:  2,
			},
			{
				NoConsume: []ResourceType{
					Human,
				},
				Consume: []ResourceType{
					Wood,
					Wood,
					Stone,
				},
				Result: Axe,
				Amount: 1,
			},
			{
				NoConsume: []ResourceType{
					Human,
				},
				Consume: []ResourceType{
					Wood,
					Stone,
					Stone,
				},
				Result: Peak,
				Amount: 1,
			},
			{
				NoConsume: []ResourceType{
					Human,
				},
				Consume: []ResourceType{
					Stone,
					Stone,
					Stone,
				},
				Result: Casserole,
				Amount: 1,
			},
			{
				NoConsume: []ResourceType{
					Human,
					Casserole,
				},
				Consume: []ResourceType{
					Food,
					Wood,
				},
				Result: Food,
				Amount: 4,
			},
			{
				NoConsume: []ResourceType{
					Human,
					Casserole,
				},
				Consume: []ResourceType{
					Human,
					Wood,
				},
				Result: Food,
				Amount: 1,
			},
		},
		Tree: {
			{
				NoConsume: []ResourceType{
					Human,
				},
				Consume: []ResourceType{},
				Result:  Wood,
				Amount:  1,
			},
			{
				NoConsume: []ResourceType{
					Human,
					Axe,
				},
				Consume: []ResourceType{},
				Result:  Wood,
				Amount:  2,
			},
			{
				NoConsume: []ResourceType{
					Human,
					Human,
				},
				Consume: []ResourceType{},
				Result:  Food,
				Amount:  2,
			},
		},
		Rock: {
			{
				NoConsume: []ResourceType{
					Human,
				},
				Consume: []ResourceType{},
				Result:  Stone,
				Amount:  1,
			},
			{
				NoConsume: []ResourceType{
					Human,
					Peak,
				},
				Consume: []ResourceType{},
				Result:  Stone,
				Amount:  2,
			},
			{
				NoConsume: []ResourceType{
					Human,
					Human,
				},
				Consume: []ResourceType{},
				Result:  Brick,
				Amount:  1,
			},
		},
	}

	recipes := RecipeDb{}

	for _, unsortedRecipeList := range unsortedRecipes {
		recipeList := make([]*Recipe, len(unsortedRecipeList))
		copy(recipeList, unsortedRecipeList)
		sort.Slice(recipeList, func(i, j int) bool {
			return len(recipeList[i].Consume)+len(recipeList[i].NoConsume) >
				len(recipeList[j].Consume)+len(recipeList[j].NoConsume)
		})
	}

	return recipes, unsortedRecipes
}

func (recipes RecipeDb) GetAllIngredientsRecipes(structureType StructureType, resources []ResourceType) (*Recipe, []ResourceType) {
	structureTypeWithHouses := structureType
	if structureType == HouseLv1 || structureType == HouseLv2 || structureType == HouseLv3 {
		structureTypeWithHouses = AllHouses
	}
	// get recipes for structure
	recipesForStructure := recipes[structureTypeWithHouses]

	// check all recipes
	selectedRecipeLen := 0
	for _, recipe := range recipesForStructure {
		var allIngredients []ResourceType = nil
		allIngredients = append(allIngredients, recipe.Consume...)
		allIngredients = append(allIngredients, recipe.NoConsume...)
		isContained, notUsed := ContainsAll(allIngredients, resources)
		if isContained &&
			selectedRecipeLen < len(allIngredients) {
			return recipe, notUsed
		}

	}

	return nil, nil
}

func ContainsAll(ingredients []ResourceType, resources []ResourceType) (bool, []ResourceType) {
	var ingredientMap map[ResourceType]int = make(map[ResourceType]int)
	var resourceMap map[ResourceType]int = make(map[ResourceType]int)

	for _, ingredient := range ingredients {
		amount, ok := ingredientMap[ingredient]
		if !ok {
			amount = 0
		}
		ingredientMap[ingredient] = amount + 1
	}

	for _, ingredient := range resources {
		amount, ok := resourceMap[ingredient]
		if !ok {
			amount = 0
		}
		resourceMap[ingredient] = amount + 1
	}

	for ingredient, amount := range ingredientMap {
		amount2, ok := resourceMap[ingredient]
		if !ok || amount2 < amount {
			return false, nil
		}
		resourceMap[ingredient] -= amount
	}

	// get not used ingredients
	notUsed := []ResourceType{}

	for ingredient, amount := range resourceMap {
		for i := 0; i < amount; i++ {
			notUsed = append(notUsed, ingredient)
		}
	}

	return true, notUsed
}
