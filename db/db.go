package db

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kreid01/gym/schema"
)

var db *gorm.DB
var err error

func InitPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USER")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(schema.Workout{})
	db.AutoMigrate(schema.Exercise{})
	db.AutoMigrate(schema.WorkoutExercise{})
	db.AutoMigrate(schema.ExerciseSet{})
	db.AutoMigrate(schema.CompletedExerciseSet{})
}

func CreateWorkout(workout *schema.Workout) (*schema.Workout, error) {
	workout.ID = uuid.New().String()
	res := db.Create(&workout)
	if res.Error != nil {
		return nil, res.Error
	}
	return workout, nil
}

func GetWorkout(id string) (*schema.Workout, error) {
	var workout schema.Workout
	res := db.First(&workout, "id = ?", id).Preload("Exercises").Preload("Exercises.Exercise").
		Find(&workout)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf(fmt.Sprintf("workout of id %s not found", id))
	}
	return &workout, nil
}

func DeleteWorkout(id string) error {
	var deletedWorkout schema.Workout
	result := db.Where("id = ?", id).Delete(&deletedWorkout)
	if result.RowsAffected == 0 {
		return fmt.Errorf("workout not deleted")
	}
	return nil
}

func GetWorkouts() ([]*schema.Workout, error) {
	var workouts []*schema.Workout
	res := db.Model(&schema.Workout{}).Preload("Exercises").Preload("Exercises.Exercise").
		Find(&workouts)

	if res.Error != nil {
		return nil, fmt.Errorf("no workouts found")
	}
	return workouts, nil
}

func UpdateWorkout(workout *schema.Workout) (*schema.Workout, error) {
	var workoutToUpdate schema.Workout
	result := db.Model(&workoutToUpdate).Where("id = ?", workout.ID).Updates(workout)
	if result.RowsAffected == 0 {
		return &workoutToUpdate, fmt.Errorf("workout not updated")
	}
	return workout, nil
}

func CreateExercise(exercise *schema.Exercise) (*schema.Exercise, error) {
	exercise.ID = uuid.New().String()
	res := db.Create(&exercise)
	if res.Error != nil {
		return nil, res.Error
	}
	return exercise, nil
}

func GetExercise(id string) (*schema.Exercise, error) {
	var exercise schema.Exercise
	res := db.First(&exercise, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("exercise of id %s not found", id)
	}
	return &exercise, nil
}

func DeleteExercise(id string) error {
	var deletedExercise schema.Exercise
	result := db.Where("id = ?", id).Delete(&deletedExercise)
	if result.RowsAffected == 0 {
		return fmt.Errorf("exercise not deleted")
	}
	return nil
}

func GetExercises() ([]*schema.Exercise, error) {
	var exercises []*schema.Exercise
	res := db.Find(&exercises)
	if res.Error != nil {
		return nil, fmt.Errorf("no exercises found")
	}
	return exercises, nil
}

func UpdateExercise(exercise *schema.Exercise) (*schema.Exercise, error) {
	var exerciseToUpdate schema.Exercise
	result := db.Model(&exerciseToUpdate).Where("id = ?", exercise.ID).Updates(exercise)
	if result.RowsAffected == 0 {
		return &exerciseToUpdate, fmt.Errorf("exercise not updated")
	}
	return exercise, nil
}

func CreateWorkoutExercise(workoutExercise *schema.WorkoutExercise) (*schema.WorkoutExercise, error) {
	workoutExercise.ID = uuid.New().String()
	res := db.Create(&workoutExercise)
	if res.Error != nil {
		return nil, res.Error
	}
	return workoutExercise, nil
}

func GetWorkoutExercise(id string) (*schema.WorkoutExercise, error) {
	var workoutExercise schema.WorkoutExercise
	res := db.First(&workoutExercise, "id = ?", id).Preload("Exercises").Preload("Sets").
		Find(&workoutExercise)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("workout exercise with id %s not found", id)
	}
	return &workoutExercise, nil
}

func DeleteWorkoutExercise(id string) error {
	var deletedWorkoutExercise schema.WorkoutExercise
	result := db.Where("id = ?", id).Delete(&deletedWorkoutExercise)
	if result.RowsAffected == 0 {
		return fmt.Errorf("workout exercise not deleted")
	}
	return nil
}

func GetWorkoutExercises() ([]*schema.WorkoutExercise, error) {
	var workoutExercises []*schema.WorkoutExercise
	res := db.Find(&workoutExercises).Preload("Exercise").Preload("Sets").
		Find(&workoutExercises)
	if res.Error != nil {
		return nil, fmt.Errorf("no workout exercises found")
	}
	return workoutExercises, nil
}

func UpdateWorkoutExercise(workoutExercise *schema.WorkoutExercise) (*schema.WorkoutExercise, error) {
	var workoutExerciseToUpdate schema.WorkoutExercise
	result := db.Model(&workoutExerciseToUpdate).Where("id = ?", workoutExercise.ID).Updates(workoutExercise)
	if result.RowsAffected == 0 {
		return &workoutExerciseToUpdate, fmt.Errorf("workout exercise not updated")
	}
	return workoutExercise, nil
}

func CreateExerciseSet(exerciseSet *schema.ExerciseSet) (*schema.ExerciseSet, error) {
	exerciseSet.ID = uuid.New().String()
	res := db.Create(&exerciseSet)
	if res.Error != nil {
		return nil, res.Error
	}
	return exerciseSet, nil
}

func GetExerciseSet(id string) (*schema.ExerciseSet, error) {
	var exerciseSet schema.ExerciseSet
	res := db.First(&exerciseSet, "id = ?", id).
		Find(&exerciseSet)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("exercise set of id %s not found", id)
	}
	return &exerciseSet, nil
}

func DeleteExerciseSet(id string) error {
	var deletedExerciseSet schema.ExerciseSet
	result := db.Where("id = ?", id).Delete(&deletedExerciseSet)
	if result.RowsAffected == 0 {
		return fmt.Errorf("exercise set not deleted")
	}
	return nil
}

func GetExerciseSets() ([]*schema.ExerciseSet, error) {
	var exerciseSets []*schema.ExerciseSet
	res := db.Model(&schema.ExerciseSet{}).
		Find(&exerciseSets)

	if res.Error != nil {
		return nil, fmt.Errorf("no exercise sets found")
	}
	return exerciseSets, nil
}

func UpdateExerciseSet(exerciseSet *schema.ExerciseSet) (*schema.ExerciseSet, error) {
	var exerciseSetToUpdate schema.ExerciseSet
	result := db.Model(&exerciseSetToUpdate).Where("id = ?", exerciseSet.ID).Updates(exerciseSet)
	if result.RowsAffected == 0 {
		return &exerciseSetToUpdate, fmt.Errorf("exercise set not updated")
	}
	return exerciseSet, nil
}

func CreateCompletedExerciseSet(completedExerciseSet *schema.CompletedExerciseSet) (*schema.CompletedExerciseSet, error) {
	completedExerciseSet.ID = uuid.New().String()
	res := db.Create(&completedExerciseSet)
	if res.Error != nil {
		return nil, res.Error
	}
	return completedExerciseSet, nil
}

func GetCompletedExerciseSet(id string) (*schema.CompletedExerciseSet, error) {
	var completedExerciseSet schema.CompletedExerciseSet
	res := db.First(&completedExerciseSet, "id = ?", id).Preload("ExerciseSet").Find(&completedExerciseSet)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("completed exercise set of id %s not found", id)
	}
	return &completedExerciseSet, nil
}

func DeleteCompletedExerciseSet(id string) error {
	var deletedCompletedExerciseSet schema.CompletedExerciseSet
	result := db.Where("id = ?", id).Delete(&deletedCompletedExerciseSet)
	if result.RowsAffected == 0 {
		return fmt.Errorf("completed exercise set not deleted")
	}
	return nil
}

func GetCompletedExerciseSets() ([]*schema.CompletedExerciseSet, error) {
	var completedExerciseSets []*schema.CompletedExerciseSet
	res := db.Model(&schema.CompletedExerciseSet{}).Preload("ExerciseSet").Find(&completedExerciseSets)

	if res.Error != nil {
		return nil, fmt.Errorf("no completed exercise sets found")
	}
	return completedExerciseSets, nil
}

func UpdateCompletedExerciseSet(completedExerciseSet *schema.CompletedExerciseSet) (*schema.CompletedExerciseSet, error) {
	var completedExerciseSetToUpdate schema.CompletedExerciseSet
	result := db.Model(&completedExerciseSetToUpdate).Where("id = ?", completedExerciseSet.ID).Updates(completedExerciseSet)
	if result.RowsAffected == 0 {
		return &completedExerciseSetToUpdate, fmt.Errorf("completed exercise set not updated")
	}
	return completedExerciseSet, nil
}
