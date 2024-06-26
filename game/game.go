package game

import (
	"log"

	"github.com/MaximoGCH/space-colony-game/assets"
	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
	"github.com/MaximoGCH/space-colony-game/game/scenes"
	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	GlobalState *state.GlobalState
}

func (g *Game) Update() error {
	switch g.GlobalState.Scene {
	case state.GameScene:
		scenes.UpdateGameScene(g.GlobalState)
	case state.MenuScene:
		scenes.UpdateMenuScene(g.GlobalState)
	case state.LostScene:
		scenes.UpdateLostScene(g.GlobalState)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.GlobalState.Scene {
	case state.GameScene:
		scenes.DrawGameScene(g.GlobalState, screen)
	case state.MenuScene:
		scenes.DrawMenuScene(g.GlobalState, screen)
	case state.LostScene:
		scenes.DrawLostScene(g.GlobalState, screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.GlobalState.ScreenSize.Width, g.GlobalState.ScreenSize.Height
}

func Start() {
	assets := assets.Load()
	globalState := &state.GlobalState{
		Assets: assets,
		ScreenSize: shapes.Size{
			Width: 912, Height: 513,
		},
		Db: db.CreateDb(assets),
	}

	scenes.StartMenuScene(globalState)

	ebiten.SetWindowSize(1480, 820)
	ebiten.SetWindowTitle("space-colony-game")
	if err := ebiten.RunGame(&Game{
		GlobalState: globalState,
	}); err != nil {
		log.Fatal(err)
	}
}
