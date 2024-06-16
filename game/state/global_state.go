package state

import (
	"github.com/MaximoGCH/space-colony-game/assets"
	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
)

type GlobalState struct {
	Scene      Scene
	Assets     *assets.Assets
	ScreenSize shapes.Size
	Db         *db.Db
	GameState  *GameState
}
