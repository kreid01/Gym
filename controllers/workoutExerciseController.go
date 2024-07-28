package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kreid01/gym/db"
	schema "github.com/kreid01/gym/schema"
)

func PostWorkoutExercise(ctx *gin.Context) {
	var workoutExercise schema.WorkoutExercise
	err := ctx.Bind(&workoutExercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateWorkoutExercise(&workoutExercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"workoutExercise": res,
	})
}

func GetWorkoutExercises(ctx *gin.Context) {
	res, err := db.GetWorkoutExercises()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"workoutExercises": res,
	})
}

func GetWorkoutExercise(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetWorkoutExercise(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"workoutExercise": res,
	})
}

func UpdateWorkoutExercise(ctx *gin.Context) {
	var updatedWorkoutExercise schema.WorkoutExercise
	err := ctx.Bind(&updatedWorkoutExercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbWorkoutExercise, err := db.GetWorkoutExercise(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbWorkoutExercise.Position = updatedWorkoutExercise.Position
	dbWorkoutExercise.Sets = updatedWorkoutExercise.Sets

	res, err := db.UpdateWorkoutExercise(dbWorkoutExercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"workoutExercise": res,
	})
}

func DeleteWorkoutExercise(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteWorkoutExercise(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "workout exercise deleted successfully",
	})
}
