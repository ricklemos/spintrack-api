package main

import (
	"fmt"
	"github.com/spintrack-api/database"
	"github.com/spintrack-api/routes"
)

func main() {
	fmt.Println("Hello SpinTrack API")
	database.ConnectToDb()
	routes.HandleRequests()
}
