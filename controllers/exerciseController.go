package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kreid01/gym/db"
	schema "github.com/kreid01/gym/schema"
)

func PostExercise(ctx *gin.Context) {
	var exercise schema.Exercise
	err := ctx.Bind(&exercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateExercise(&exercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"exercise": res,
	})
}

func GetExercises(ctx *gin.Context) {
	res, err := db.GetExercises()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"exercises": res,
	})
}

func GetExercise(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetExercise(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"exercise": res,
	})
}

func UpdateExercise(ctx *gin.Context) {
	var updatedExercise schema.Exercise
	err := ctx.Bind(&updatedExercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbExercise, err := db.GetExercise(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbExercise.Name = updatedExercise.Name
	dbExercise.MuscleGroup = updatedExercise.MuscleGroup

	res, err := db.UpdateExercise(dbExercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"exercise": res,
	})
}

func DeleteExercise(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteExercise(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "exercise deleted successfully",
	})
}
