package systems

import (
	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateNextDay(globalState *state.GlobalState) {
	UpdateButton(globalState.GameState.NextDayButton)

	if globalState.GameState.NextDayButton.IsJustActive &&
		globalState.GameState.NextDayTransitionPhase == 0 {
		globalState.GameState.NextDayTransitionPhase = 1

		// set next day state
		globalState.GameState.NextDayState = &models.NextDay{
			Timer: 0,
		}
	}

	if globalState.GameState.NextDayTransitionPhase > 0 {
		globalState.GameState.NextDayState.Timer++
	}

	switch globalState.GameState.NextDayTransitionPhase {
	case 1:
		// wait for all lost cards
		if len(globalState.GameState.LostResourceCard) == 0 &&
			globalState.GameState.NextDayState.Timer > 50 {
			globalState.GameState.NextDayTransitionPhase++
			globalState.GameState.NextDayState.Timer = 0
		}

	case 2:
		// check all recipes
		for i, boardLine := range globalState.GameState.Board {
			for j, boardStructure := range boardLine {
				timer := (i + j*3) * 10
				println(timer)
				if boardStructure == nil || timer != globalState.GameState.NextDayState.Timer {
					continue
				}

				for _, group := range boardStructure.CardDrop {
					groupIngredients := []db.ResourceType{}
					for _, cardDrop := range group {
						if cardDrop.Type == db.Empty {
							continue
						}

						groupIngredients = append(groupIngredients, cardDrop.Type)
						cardDrop.Type = db.Empty
					}

					if len(groupIngredients) == 0 {
						continue
					}

					// check recipes
					selectedRecipe, resourcesToReturn := globalState.Db.Recipes.GetAllIngredientsRecipes(boardStructure.Type, groupIngredients)

					if selectedRecipe == nil {
						// no card was used
						resourcesToReturn = groupIngredients
					} else {
						// add not consumed cards
						resourcesToReturn = append(resourcesToReturn, selectedRecipe.NoConsume...)

						// add recipe return
						for i := 0; i < selectedRecipe.Amount; i++ {
							resourcesToReturn = append(resourcesToReturn, selectedRecipe.Result)
						}
					}

					// generate lost cards for every resource new and not used or not consumed
					for _, resource := range resourcesToReturn {
						globalState.GameState.LostResourceCard.Add(
							&models.ResourceCard{
								Type:        resource,
								Position:    boardStructure.Position.ConstAdd(models.GridSize / 2),
								MouseOffset: shapes.Point{},
							},
							globalState.GameState.Inventory,
							true,
						)
					}
				}
			}
		}

		maxTimer := 10 * (models.BoardSizeH*models.BoardSizeW + 1)
		if globalState.GameState.NextDayState.Timer == maxTimer &&
			len(globalState.GameState.LostResourceCard) == 0 {
			globalState.GameState.NextDayTransitionPhase++
		}
	}
}

func DrawNextDay(globalState *state.GlobalState, screen *ebiten.Image) {
	DrawButton(globalState, screen, globalState.GameState.NextDayButton)
}
