package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spintrack-api/controllers"
	"log"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/api/activities", controllers.GetActivities)

	err := r.Run()
	if err != nil {
		log.Panic("Error running Gin", err.Error())
	}
}
