package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kreid01/gym/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/workouts", controllers.GetWorkouts)
	r.GET("/workouts/:id", controllers.GetWorkout)
	r.POST("/workouts", controllers.PostWorkout)
	r.PUT("/workouts/:id", controllers.UpdateWorkout)
	r.DELETE("/workouts/:id", controllers.DeleteWorkout)

	r.GET("/exercise", controllers.GetExercises)
	r.GET("/exercise/:id", controllers.GetExercise)
	r.POST("/exercise", controllers.PostExercise)
	r.PUT("/exercise/:id", controllers.UpdateExercise)
	r.DELETE("/exercise/:id", controllers.DeleteExercise)

	r.GET("/workoutExercises", controllers.GetWorkoutExercises)
	r.GET("/workoutExercise/:id", controllers.GetWorkoutExercise)
	r.POST("/workoutExercise", controllers.PostWorkoutExercise)
	r.PUT("/workoutExercise/:id", controllers.UpdateWorkoutExercise)
	r.DELETE("/workoutExercise/:id", controllers.DeleteWorkoutExercise)
	return r
}
