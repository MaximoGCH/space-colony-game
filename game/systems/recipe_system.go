package systems

import (
	"fmt"
	"image/color"

	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/custom_text"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	RecipeXStart = 8
	RecipeYStart = 16
)

func UpdateRecipes(globalState *state.GlobalState) {
	_, dy := ebiten.Wheel()
	globalState.GameState.RecipeListControls.Scroll += int(dy * 20)
	mousePos := shapes.FromMousePosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) && mousePos.X < 196 {
		globalState.GameState.RecipeListControls.MouseOffset =
			globalState.GameState.RecipeListControls.Scroll - mousePos.Y
		globalState.GameState.RecipeListControls.IsMouseScrolling = true
	}

	if globalState.GameState.RecipeListControls.IsMouseScrolling {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
			globalState.GameState.RecipeListControls.IsMouseScrolling = false
		}

		globalState.GameState.RecipeListControls.Scroll =
			mousePos.Y + globalState.GameState.RecipeListControls.MouseOffset
	}
}

func DrawRecipes(globalState *state.GlobalState, screen *ebiten.Image) {
	position := shapes.Point{
		X: RecipeXStart,
		Y: RecipeYStart + globalState.GameState.RecipeListControls.Scroll,
	}

	rec := ebiten.NewImage(196, globalState.ScreenSize.Height)
	rec.Fill(color.Black)
	screen.DrawImage(rec, shapes.Point{}.ToImageOptions())

	cardImage := globalState.Assets.GetSprite("embed/sprites/card")

	for _, structureType := range globalState.Db.StructureList {
		recipeList := globalState.Db.UnsortedRecipes[structureType]
		structureTypeWithHouse := structureType

		if structureTypeWithHouse == db.AllHouses {
			structureTypeWithHouse = db.HouseLv1
		}

		structureImage := globalState.Db.Structures[structureTypeWithHouse].Image

		screen.DrawImage(structureImage, position.PointAdd(shapes.Point{X: 56, Y: 0}).ToImageOptions())

		position = position.PointAdd(shapes.Point{X: 0, Y: 80})

		for _, recipe := range recipeList {
			ingredients := []db.ResourceType{}
			ingredients = append(ingredients, recipe.NoConsume...)
			ingredients = append(ingredients, recipe.Consume...)

			for i, ingredient := range ingredients {
				ingredientImage := globalState.Db.Resources[ingredient].BigImage
				screen.DrawImage(ingredientImage, position.PointAdd(shapes.Point{
					X: i * 32, Y: 0,
				}).ToImageOptions())
			}

			resultPoint := position.PointAdd(shapes.Point{
				X: 4*32 + 8, Y: 0,
			})
			resultImage := globalState.Db.Resources[recipe.Result].BigImage
			screen.DrawImage(cardImage, resultPoint.ToImageOptions())
			screen.DrawImage(resultImage, resultPoint.ToImageOptions())
			custom_text.DrawOutlineText(screen,
				fmt.Sprintf("x%v", recipe.Amount),
				resultPoint.PointAdd(shapes.Point{X: 16, Y: 38}),
				globalState.Assets.GetFont("embed/fonts/Kubasta"),
			)

			position = position.PointAdd(shapes.Point{X: 0, Y: 64})
		}

		position = position.PointAdd(shapes.Point{X: 0, Y: 32})
	}
}
