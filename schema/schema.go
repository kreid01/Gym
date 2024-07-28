package schema

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
}

type Workout struct {
	ID        string            `json:"id" gorm:"primaryKey"`
	Name      string            `json:"name"`
	User      User              `json:"user" gorm:"foreignKey:UserRefer"`
	Exercises []WorkoutExercise `json:"exercises" gorm:"foreignKey:WorkoutRefer"`
	UserRefer string
}

type Exercise struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	MuscleGroup string `json:"muscleGroup"`
}

type WorkoutExercise struct {
	ID            string        `json:"id" gorm:"primaryKey"`
	WorkoutRefer  string        `json:"workoutRefer"`
	ExerciseRefer string        `json:"exerciseRefer"`
	Exercise      Exercise      `json:"exercise" gorm:"foreignKey:ExerciseRefer"`
	Position      int           `json:"position"`
	Sets          []ExerciseSet `json:"sets" gorm:"foreignKey:WorkoutExerciseRefer"`
}

type ExerciseSet struct {
	ID                   string `json:"id" gorm:"primaryKey"`
	WorkoutExerciseRefer string
	Weight               string `json:"weight"`
	Repetitions          string `json:"repetitions"`
}

type CompletedExerciseSet struct {
	ID               string `json:"id" gorm:"primaryKey"`
	ExerciseSetRefer string
	ExerciseSet      ExerciseSet `gorm:"foreignKey:ExerciseSetRefer"`
	Repetitions      string      `json:"maxReps"`
}
