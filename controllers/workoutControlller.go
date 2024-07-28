package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kreid01/gym/db"
	schema "github.com/kreid01/gym/schema"
)

func PostWorkout(ctx *gin.Context) {
	var workout schema.Workout
	err := ctx.Bind(&workout)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateWorkout(&workout)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"workout": res,
	})
}

func GetWorkouts(ctx *gin.Context) {
	res, err := db.GetWorkouts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"workout": res,
	})
}

func GetWorkout(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetWorkout(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"workout": res,
	})
}

func UpdateWorkout(ctx *gin.Context) {
	var updatedWorkout schema.Workout
	err := ctx.Bind(&updatedWorkout)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbWorkout, err := db.GetWorkout(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbWorkout.Name = updatedWorkout.Name

	res, err := db.UpdateWorkout(dbWorkout)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"task": res,
	})
}

func DeleteWorkout(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteWorkout(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}
