package types

type WorkoutViewModel struct {
	ID           string `json:"id"`
	KcalBurned   int    `json:"kcalBurned"`
	ActivityType string `json:"activityType"`
	Name         string `json:"name"`
	Duration     string `json:"duration"`
	Picto        string `json:"picto"`
}

type WorkoutInputModel struct {
	KcalBurned   int     `json:"kcalBurned"`
	ActivityType string  `json:"activityType"`
	Name         string  `json:"name"`
	Duration     float64 `json:"duration"`
}

type HevyRes struct {
	Page      int           `json:"page"`
	PageCount int           `json:"pageCount"`
	Workouts  []HevyWorkout `json:"workouts"`
}

type HevyWorkout struct {
	ID        string         `json:"id"`
	Title     string         `json:"title"`
	Exercices []HevyExercice `json:"exercices"`
	CreatedAt string         `json:"createdAt"`
}

type HevyExercice struct {
	Sets []HevySets `json:"sets"`
}

type HevySets struct {
	SetType string `json:"setType"`
}
