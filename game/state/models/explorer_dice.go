package models

type Dice struct {
	FaceNumber  int
	Bounce      int
	BounceTimer int
}

type ExplorerDices [ExplorerCardDropLen]*Dice

func CreateExplorerDices() ExplorerDices {
	return ExplorerDices{}
}
