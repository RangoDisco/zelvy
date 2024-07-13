package models

type GoalType int

const (
	Steps GoalType = iota
	KcalBurned
	KcalConsumed
	WorkoutDuration
)

func (g GoalType) String() string {
	return [...]string{"KcalBurned", "KcalConsumed", "WorkoutDuration"}[g]
}
