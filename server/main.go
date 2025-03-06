package main

import (
	"os"

	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	// these are the endpoints
	//C
	router.POST("/task/create", routes.CreateTask)
	//R
	router.GET("/task/get", routes.GetTask)
	//U
	router.PUT("/task/update/:id", routes.UpdateTask)
	//D
	router.DELETE("/task/delete/:id", routes.DeleteTask)

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}