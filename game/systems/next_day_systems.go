package systems

import (
	"fmt"

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
		// set next day state
		fixedDay := globalState.GameState.Days
		if fixedDay >= 6 {
			fixedDay = 6
		}
		globalState.GameState.NextDayState = &models.NextDay{
			Timer:    0,
			FixedDay: fixedDay,
		}

		globalState.GameState.NextDayTransitionPhase = 1
	}

	if globalState.GameState.NextDayTransitionPhase > 0 {
		globalState.GameState.NextDayState.Timer++
	}

	switch globalState.GameState.NextDayTransitionPhase {
	case 1:
		if globalState.GameState.NextDayState.Timer == 30 {
			globalState.GameState.Notifications.Add(models.Text,
				"Work time")
		}

		// wait for all lost cards
		if globalState.GameState.NextDayState.Timer == 50 {
			NextState(globalState, false)
		}

	case 2:
		// check all recipes
		for i := 0; i < len(globalState.GameState.Board); i++ {
			for j := 0; j < len(globalState.GameState.Board[i]); j++ {
				boardStructure := globalState.GameState.Board[i][j]
				checkedPos := j + i*len(globalState.GameState.Board[i])

				if boardStructure == nil || globalState.GameState.NextDayState.Timer <
					20 || globalState.GameState.NextDayState.BoardCheckedPos >= checkedPos {
					continue
				}

				globalState.GameState.NextDayState.BoardCheckedPos = checkedPos

				hasIngredients := false
				removeStructure := false
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

					hasIngredients = true

					// check recipes
					selectedRecipe, resourcesToReturn := globalState.Db.Recipes.GetAllIngredientsRecipes(boardStructure.Type, groupIngredients)

					if selectedRecipe == nil {
						// no card was used
						resourcesToReturn = groupIngredients
					} else {
						// add notification
						resourceName := globalState.Db.Resources[selectedRecipe.Result].Name
						globalState.GameState.Notifications.Add(models.Text,
							fmt.Sprintf("+%v %v", selectedRecipe.Amount, resourceName))

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

					// remove resource
					removeStructure = true
				}

				if hasIngredients {
					globalState.GameState.NextDayState.Timer = 0
				}

				if removeStructure && (boardStructure.Type != db.HouseLv1) &&
					boardStructure.Type != db.HouseLv2 && boardStructure.Type != db.HouseLv3 {
					globalState.GameState.Board[i][j] = nil
				}
			}
		}

		if globalState.GameState.NextDayState.Timer == 30 {
			WaitAndNextState(globalState)
		}

	case 3:

		if globalState.GameState.NextDayState.Timer == 50 {
			globalState.GameState.Notifications.Add(models.Text,
				"Exploration time")
		}

		if globalState.GameState.NextDayState.Timer == 80 {
			if len(globalState.GameState.Board.EmptyPlaces()) == 0 {
				globalState.GameState.Notifications.Add(models.Text,
					"You know too many resources to explore further")

				globalState.GameState.NextDayState.SkipTurn = true
			} else {
				carDropWithResources := false
				for _, cardDrop := range globalState.GameState.ExplorerCardDrop {
					if cardDrop.Type != db.Empty {
						carDropWithResources = true
						break
					}
				}

				if !carDropWithResources {
					globalState.GameState.Notifications.Add(models.Text,
						"You have not sent any explorer ")

					globalState.GameState.NextDayState.SkipTurn = true
				} else {
					if globalState.GameState.Days <= 3 {
						globalState.GameState.Notifications.Add(models.Text,
							"There are abundant resources")
					} else if globalState.GameState.Days <= 4 {
						globalState.GameState.Notifications.Add(models.Text,
							"There are a few resources")
					} else {
						globalState.GameState.Notifications.Add(models.Text,
							"Resources are scarce")
					}
				}
			}
		}

		if globalState.GameState.NextDayState.SkipTurn &&
			globalState.GameState.NextDayState.Timer == 120 {
			WaitAndGoToState(globalState, 6)
		}

		if globalState.GameState.NextDayState.Timer == 120 {
			globalState.GameState.Notifications.Add(models.Text,
				fmt.Sprintf("You need to roll at least a %v to find new resources", globalState.GameState.Days))
		}

		// wait for all lost cards
		if globalState.GameState.NextDayState.Timer == 150 {
			WaitAndNextState(globalState)
		}

	case 4:
		if globalState.GameState.NextDayState.Timer == 40 {
			for i, cardDrop := range globalState.GameState.ExplorerCardDrop {
				if cardDrop.Type != db.Human {
					continue
				}

				globalState.GameState.ExplorerDices[i] = &models.Dice{
					FaceNumber:  1,
					Bounce:      0,
					BounceTimer: 0,
				}
			}
		}

		if globalState.GameState.NextDayState.Timer == 130 {
			NextState(globalState, false)
		}

	case 5:
		if globalState.GameState.NextDayState.Timer == 10 {
			globalState.GameState.NextDayState.DiceResult = false
			for _, dice := range globalState.GameState.ExplorerDices {
				if dice == nil {
					continue
				}

				if dice.FaceNumber >= globalState.GameState.Days {
					globalState.GameState.NextDayState.DiceResult = true
					break
				}
			}
		}

		if globalState.GameState.NextDayState.Timer == 50 {
			if globalState.GameState.NextDayState.DiceResult {
				globalState.GameState.Notifications.Add(models.Text, "You have found new resources")
			} else {
				globalState.GameState.Notifications.Add(models.Text, "You have not found new resources")
			}
		}

		if globalState.GameState.NextDayState.DiceResult {
			if globalState.GameState.NextDayState.Timer == 100 {
				globalState.GameState.Board.AddStructure(globalState.ScreenSize,
					globalState.Db.Structures[db.Tree], globalState.GameState.Board.RandomEmptyPlace())
			}

			if globalState.GameState.NextDayState.Timer == 130 {
				globalState.GameState.Board.AddStructure(globalState.ScreenSize,
					globalState.Db.Structures[db.Rock], globalState.GameState.Board.RandomEmptyPlace())
			}

			if globalState.GameState.NextDayState.Timer == 160 {
				globalState.GameState.Board.AddStructure(globalState.ScreenSize,
					globalState.Db.Structures[db.Tree], globalState.GameState.Board.RandomEmptyPlace())
			}
		}

		if globalState.GameState.NextDayState.Timer == 260 {
			for i, cardDrop := range globalState.GameState.ExplorerCardDrop {
				if cardDrop.Type == db.Empty {
					continue
				}

				globalState.GameState.ExplorerDices[i] = nil

				globalState.GameState.LostResourceCard.Add(
					&models.ResourceCard{
						Type:        cardDrop.Type,
						Position:    cardDrop.Bounds.Center(),
						MouseOffset: shapes.Point{},
					},
					globalState.GameState.Inventory,
					true,
				)

				cardDrop.Type = db.Empty
			}
		}

		if globalState.GameState.NextDayState.Timer == 350 {
			WaitAndNextState(globalState)
		}

	case 6:
		numberHumans := globalState.GameState.Inventory[db.Human].Amount
		numberFood := globalState.GameState.Inventory[db.Food].Amount

		if globalState.GameState.NextDayState.Timer == 50 {
			globalState.GameState.Notifications.Add(models.Text,
				fmt.Sprintf("You have %v humans in your colony", numberHumans))
		}

		if globalState.GameState.NextDayState.Timer == 80 {
			globalState.GameState.Notifications.Add(models.Text,
				fmt.Sprintf("Your humans need %v unit of food", numberHumans))
		}

		if globalState.GameState.NextDayState.Timer == 120 {
			if numberFood >= numberHumans {
				globalState.GameState.Notifications.Add(models.Text,
					fmt.Sprintf("Your humans eat %v unit of food", numberHumans))
				globalState.GameState.Inventory.RemoveResource(db.Food, numberHumans)
			} else {
				globalState.GameState.Notifications.Add(models.Text,
					fmt.Sprintf("Your humans eat %v unit of food", numberFood))
				globalState.GameState.NextDayState.HumanDied = numberHumans - numberFood
				globalState.GameState.Inventory.RemoveResource(db.Food, numberFood)
			}
		}

		if globalState.GameState.NextDayState.Timer == 150 {
			if globalState.GameState.NextDayState.HumanDied == 0 {
				globalState.GameState.Notifications.Add(models.Text, "All humans have survived the day")
			} else {
				globalState.GameState.Notifications.Add(models.Text,
					fmt.Sprintf("%v humans have died of starvation",
						globalState.GameState.NextDayState.HumanDied))

				globalState.GameState.Inventory.RemoveResource(db.Human,
					globalState.GameState.NextDayState.HumanDied)
			}
		}

		if globalState.GameState.NextDayState.Timer == 180 {
			WaitAndNextState(globalState)
		}

	case 7:
		for i := range globalState.GameState.ExplorerDices {
			globalState.GameState.ExplorerDices[i] = nil
		}

		if globalState.GameState.NextDayState.Timer == 50 {
			globalState.GameState.Days++
			globalState.GameState.Notifications.Add(models.Text,
				fmt.Sprintf("Day %v", globalState.GameState.Days))
		}

		if globalState.GameState.NextDayState.Timer == 80 {
			WaitAndGoToState(globalState, 0)
		}
	}
}

func DrawNextDay(globalState *state.GlobalState, screen *ebiten.Image) {
	DrawButton(globalState, screen, globalState.GameState.NextDayButton)
}

func WaitAndNextState(globalState *state.GlobalState) {
	WaitAndGoToState(globalState, globalState.GameState.NextDayTransitionPhase+1)
}

func WaitAndGoToState(globalState *state.GlobalState, state int) {
	if !globalState.GameState.NextDayState.ClearCheck {
		globalState.GameState.ClearNotifications = true
		globalState.GameState.NextDayState.ClearCheck = true
	}

	globalState.GameState.NextDayState.Timer--

	if globalState.GameState.NotificationsJustCleared {
		GoToState(globalState, state, true)
		globalState.GameState.NextDayState.ClearCheck = false
	}
}

func NextState(globalState *state.GlobalState, clear bool) {
	GoToState(globalState, globalState.GameState.NextDayTransitionPhase+1, clear)
}

func GoToState(globalState *state.GlobalState, state int, clear bool) {
	globalState.GameState.NextDayTransitionPhase = state
	globalState.GameState.NextDayState.Timer = 0
	globalState.GameState.NextDayState.SkipTurn = false

	if clear {
		globalState.GameState.Notifications.Clear()
	}
}
