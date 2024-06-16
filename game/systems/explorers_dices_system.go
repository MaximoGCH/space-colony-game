package systems

import (
	"fmt"
	"math/rand"

	"github.com/MaximoGCH/space-colony-game/game/state"
	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateExplorerDices(globalState *state.GlobalState) {
	for _, dice := range globalState.GameState.ExplorerDices {
		if dice == nil {
			continue
		}

		if dice.Bounce < 6 {
			dice.BounceTimer++
			if dice.BounceTimer > 15 {
				dice.FaceNumber = 1 + rand.Intn(6)
				dice.BounceTimer = 0
				dice.Bounce++
			}
		}
	}
}

func DrawExplorerDices(globalState *state.GlobalState, screen *ebiten.Image) {
	for i, cardDrop := range globalState.GameState.ExplorerCardDrop {
		dice := globalState.GameState.ExplorerDices[i]
		if dice == nil {
			continue
		}

		position := cardDrop.Bounds.Center()
		image := globalState.Assets.GetSprite(fmt.Sprintf("embed/sprites/dice-%v", dice.FaceNumber))
		screen.DrawImage(image, position.ConstSub(24).ToImageOptions())
	}
}
