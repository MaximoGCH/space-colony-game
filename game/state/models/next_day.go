package models

type NextDay struct {
	Timer           int
	ClearCheck      bool
	SkipTurn        bool
	DiceResult      bool
	BoardCheckedPos int
	HumanDied       int
	FixedDay        int
}
