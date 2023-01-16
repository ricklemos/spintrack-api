package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spintrack-api/database"
	"github.com/spintrack-api/models"
	"net/http"
)

func GetActivities(c *gin.Context) {
	var activities []models.Activity
	database.DB.Find(&activities)
	c.JSON(200, activities)
}

func CreateNewActivity(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&activity)
	c.JSON(http.StatusOK, activity)
}
