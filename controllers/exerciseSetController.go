package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kreid01/gym/db"
	schema "github.com/kreid01/gym/schema"
)

func PostExerciseSet(ctx *gin.Context) {
	var exerciseSet schema.ExerciseSet
	err := ctx.Bind(&exerciseSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateExerciseSet(&exerciseSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"exerciseSet": res,
	})
}

func GetExerciseSets(ctx *gin.Context) {
	res, err := db.GetExerciseSets()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"exerciseSets": res,
	})
}

func GetExerciseSet(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetExerciseSet(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"exerciseSet": res,
	})
}

func UpdateExerciseSet(ctx *gin.Context) {
	var updatedExerciseSet schema.ExerciseSet
	err := ctx.Bind(&updatedExerciseSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbExerciseSet, err := db.GetExerciseSet(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbExerciseSet.Repetitions = updatedExerciseSet.Repetitions
	dbExerciseSet.Weight = updatedExerciseSet.Weight

	res, err := db.UpdateExerciseSet(dbExerciseSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"exerciseSet": res,
	})
}

func DeleteExerciseSet(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteExerciseSet(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "exercise set deleted successfully",
	})
}
