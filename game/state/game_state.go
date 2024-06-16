package state

import "github.com/MaximoGCH/space-colony-game/game/state/models"

type GameState struct {
	ResourceCard           *models.ResourceCard
	LostResourceCard       models.LostResourceCardList
	Inventory              models.Inventory
	Board                  models.Board
	ExplorerCardDrop       models.ExplorerCardDrop
	NextDayButton          *models.Button
	NextDayTransitionPhase int
	NextDayState           *models.NextDay
}
