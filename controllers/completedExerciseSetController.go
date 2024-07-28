package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kreid01/gym/db"
	schema "github.com/kreid01/gym/schema"
)

func PostCompletedExerciseSet(ctx *gin.Context) {
	var completedExerciseSet schema.CompletedExerciseSet
	err := ctx.Bind(&completedExerciseSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateCompletedExerciseSet(&completedExerciseSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"completedExerciseSet": res,
	})
}

func GetCompletedExerciseSets(ctx *gin.Context) {
	res, err := db.GetCompletedExerciseSets()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"completedExerciseSets": res,
	})
}

func GetCompletedExerciseSet(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetCompletedExerciseSet(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"completedExerciseSet": res,
	})
}

func UpdateCompletedExerciseSet(ctx *gin.Context) {
	var updatedCompletedExerciseSet schema.CompletedExerciseSet
	err := ctx.Bind(&updatedCompletedExerciseSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbCompletedExerciseSet, err := db.GetCompletedExerciseSet(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbCompletedExerciseSet.Repetitions = updatedCompletedExerciseSet.Repetitions

	res, err := db.UpdateCompletedExerciseSet(dbCompletedExerciseSet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"completedExerciseSet": res,
	})
}

func DeleteCompletedExerciseSet(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteCompletedExerciseSet(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "completed exercise set deleted successfully",
	})
}
