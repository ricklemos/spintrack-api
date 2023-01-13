package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spintrack-api/database"
	"github.com/spintrack-api/models"
)

func GetActivities(c *gin.Context) {
	var activities []models.Activity
	database.DB.Find(&activities)
	c.JSON(200, activities)
}

func CreateNewActivity(c *gin.Context) {

}
