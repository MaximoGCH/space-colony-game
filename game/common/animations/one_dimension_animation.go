package animations

import "github.com/MaximoGCH/space-colony-game/game/common/animations/easing"

type OneDimensionAnimation struct {
	Value       int
	EasingValue int
	MinValue    int
	MaxValue    int
}

func NewOneDimensionAnimation(minValue int, maxValue int) *OneDimensionAnimation {
	return &OneDimensionAnimation{
		Value:       minValue,
		EasingValue: minValue,
		MinValue:    minValue,
		MaxValue:    maxValue,
	}
}

func (state *OneDimensionAnimation) Update(isDirectionForward bool, speed int, easingFunc easing.Easing) {
	var valueAfterState int
	currentValue := state.Value

	if isDirectionForward {
		newValue := currentValue + speed

		if newValue < state.MaxValue {
			valueAfterState = newValue
		} else {
			valueAfterState = state.MaxValue
		}
	} else {
		newValue := currentValue - speed

		if newValue > state.MinValue {
			valueAfterState = newValue
		} else {
			valueAfterState = state.MinValue
		}
	}

	state.Value = valueAfterState
	state.EasingValue = easing.ApplyEasing(valueAfterState, state.MinValue, state.MaxValue, easingFunc)
}
