package models

import (
	"math/rand"

	"github.com/MaximoGCH/space-colony-game/db"
	"github.com/MaximoGCH/space-colony-game/game/common/animations"
	"github.com/MaximoGCH/space-colony-game/game/common/shapes"
)

const (
	BoardSizeW = 5
	BoardSizeH = 3
)

type CardDrop struct {
	Type   db.ResourceType
	Bounds shapes.Rectangle
}

type BoardStructure struct {
	Type            db.StructureType
	CardDrop        [][]*CardDrop
	Position        shapes.Point
	AnimationOffset *animations.OneDimensionAnimation
}

type Board [BoardSizeW][BoardSizeH]*BoardStructure

const (
	GridSize          = 64
	GridSpace         = GridSize + 48
	CardDropSize      = CardWidth + 4
	CardDropGroupSize = CardHeight + 4
	EmptyListLen      = BoardSizeW * BoardSizeH
	BoardStartX       = 128
	BoardStartY       = 8
)

func CreateBoard() Board {
	return Board{}
}

func (board *Board) EmptyPlaces() []shapes.Point {
	var emptyList []shapes.Point = make([]shapes.Point, 0, BoardSizeW)
	for i := 1; i < BoardSizeW; i++ {
		for j := 0; j < BoardSizeH; j++ {
			structure := board[i][j]
			if structure == nil {
				emptyList = append(emptyList, shapes.Point{X: i, Y: j})
			}
		}
	}

	return emptyList
}

func (board *Board) RandomEmptyPlace() shapes.Point {
	emptyList := board.EmptyPlaces()

	if len(emptyList) == 0 {
		return shapes.Point{X: -1, Y: -1}
	}

	index := rand.Intn(len(emptyList))

	return emptyList[index]
}

func (board *Board) AddStructure(screenSize shapes.Size, structure *db.Structure, position shapes.Point) *BoardStructure {
	// create card drop list
	cardDrop := make([][]*CardDrop, structure.CardDropGroupNumber)

	structurePos := position.ConstMul(GridSpace).PointAdd(shapes.Point{
		X: BoardStartX, Y: BoardStartY,
	})

	cardDropXStartPos := (GridSize / 2) - ((structure.CardDropNumber * CardDropSize) / 2)
	for i := 0; i < structure.CardDropGroupNumber; i++ {
		cardDrop[i] = make([]*CardDrop, structure.CardDropNumber)

		for j := 0; j < structure.CardDropNumber; j++ {
			cardDropPosition := shapes.Point{
				X: structurePos.X + cardDropXStartPos + position.X + j*CardDropSize,
				Y: structurePos.Y + position.Y + GridSize + i*CardDropGroupSize,
			}

			cardDrop[i][j] = &CardDrop{
				Type: db.Empty,
				Bounds: shapes.Rectangle{
					Point: cardDropPosition,
					Size:  shapes.Size{Width: CardWidth, Height: CardHeight},
				},
			}
		}
	}

	// create structure
	newStructure := &BoardStructure{
		Type:            structure.Type,
		CardDrop:        cardDrop,
		Position:        structurePos,
		AnimationOffset: animations.NewOneDimensionAnimation(0, 500),
	}

	board[position.X][position.Y] = newStructure

	return newStructure
}

func (board *Board) RemoveStructure(position shapes.Point) {
	board[position.X][position.Y] = nil
}
