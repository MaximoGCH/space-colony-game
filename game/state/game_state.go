package state

import "github.com/MaximoGCH/space-colony-game/game/state/models"

type GameState struct {
	ResourceCard             *models.ResourceCard
	LostResourceCard         models.LostResourceCardList
	Inventory                models.Inventory
	Board                    models.Board
	ExplorerCardDrop         models.ExplorerCardDrop
	NextDayButton            *models.Button
	NextDayTransitionPhase   int
	NextDayState             *models.NextDay
	Notifications            models.NotificationSystem
	ClearNotifications       bool
	ClearNotificationsCheck  bool
	NotificationsJustCleared bool
	ExplorerDices            models.ExplorerDices
	Days                     int
	GameStarted              bool
	StartTimer               int
	RecipeListControls       *models.RecipeListControls
	GameOver                 bool

	StartButton *models.Button
}
