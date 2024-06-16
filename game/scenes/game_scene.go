package scenes

import (
	"image/color"

	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/MaximoGCH/space-colony-game/game/state/models"
	"github.com/MaximoGCH/space-colony-game/game/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	startTreeResources = 6
	startRockResources = 3
)

func StartGameScene(globalState *state.GlobalState) {
	globalState.Scene = state.GameScene

	// create game state
	globalState.GameState = &state.GameState{
		ResourceCard: nil,
		Inventory: models.CreateInventory(
			globalState.Db.Resources,
			globalState.Db.ResourceList,
			globalState.ScreenSize,
		),
		Board:            models.CreateBoard(),
		LostResourceCard: nil,
		ExplorerCardDrop: models.CreateExplorerCardDrop(),
		NextDayButton: models.CreateButton(shapes.Rectangle{
			Point: shapes.Point{X: globalState.ScreenSize.Width - 80, Y: globalState.ScreenSize.Height - 50},
			Size:  shapes.Size{Width: 64, Height: 32},
		}, "Next day"),
		ExplorerDices: models.CreateExplorerDices(),
		GameStarted:   false,
		RecipeListControls: &models.RecipeListControls{
			Scroll:      0,
			MouseOffset: 0,
		},
		Days: 1,
	}

	for i := 0; i < startRockResources; i++ {
		randPos := globalState.GameState.Board.RandomEmptyPlace()

		if randPos.X == -1 {
			break
		}

		structure := globalState.GameState.Board.AddStructure(
			globalState.ScreenSize,
			globalState.Db.Structures[db.Rock],
			randPos,
		)

		// no animation for first resources
		structure.AnimationOffset.Value = structure.AnimationOffset.MaxValue
	}

	for i := 0; i < startTreeResources; i++ {
		randPos := globalState.GameState.Board.RandomEmptyPlace()

		if randPos.X == -1 {
			break
		}

		structure := globalState.GameState.Board.AddStructure(
			globalState.ScreenSize,
			globalState.Db.Structures[db.Tree],
			randPos,
		)

		// no animation for first resources
		structure.AnimationOffset.Value = structure.AnimationOffset.MaxValue
	}
}

func UpdateGameScene(globalState *state.GlobalState) {
	if !globalState.GameState.GameStarted {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			globalState.GameState.GameStarted = true

			// add spaceship
			globalState.GameState.Board.AddStructure(
				globalState.ScreenSize,
				globalState.Db.Structures[db.HouseLv1],
				shapes.Point{X: 0, Y: 1},
			)

			globalState.GameState.Notifications.Clear()
		}

		if globalState.GameState.StartTimer == 30 {
			globalState.GameState.Notifications.Add(models.Text, "Welcome to planet Ionia")
		}

		if globalState.GameState.StartTimer == 40 {
			globalState.GameState.Notifications.Add(models.Text,
				"Manage your colony to last the maximum number of days.")
		}

		if globalState.GameState.StartTimer == 50 {
			globalState.GameState.Notifications.Add(models.Text,
				"Every day there are fewer resources, manage resources well because they are scarce..")
		}

		if globalState.GameState.StartTimer == 60 {
			globalState.GameState.Notifications.Add(models.Text,
				"Click to start your colony")
		}

		globalState.GameState.StartTimer++
	}

	systems.UpdateRecipes(globalState)
	systems.UpdateBoard(globalState)
	systems.UpdateExplorerCardDrop(globalState)
	systems.UpdateInventory(globalState)
	systems.UpdateResourceCard(globalState)
	systems.UpdateLostCardSystem(globalState)
	systems.UpdateNextDay(globalState)
	systems.UpdateNotifications(globalState)
	systems.UpdateExplorerDices(globalState)

	if globalState.GameState.GameOver {
		StartLostScene(globalState, globalState.GameState.Days)
	}
}

func DrawGameScene(globalState *state.GlobalState, screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 203,
		G: 129,
		B: 117,
		A: 255,
	})
	systems.DrawRecipes(globalState, screen)
	systems.DrawBoard(globalState, screen)
	systems.DrawExplorerCardDrop(globalState, screen)
	systems.DrawInventory(globalState, screen)
	systems.DrawResourceCard(globalState, screen)
	systems.DrawLostCardCard(globalState, screen)
	systems.DrawNextDay(globalState, screen)
	systems.DrawNotifications(globalState, screen)
	systems.DrawExplorerDices(globalState, screen)
}
