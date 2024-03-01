package main

import (
	"fmt"

	"project3/dbhandler"
	"project3/middleware"
	"project3/router"

	"github.com/gin-gonic/gin"
)

func handlePanic() {
	a := recover()
	if a != nil {
		fmt.Println("RECOVER", a)
	}
}

func main() {
	defer handlePanic()              // Recover if panic occurs
	defer dbhandler.DbClient.Close() //Close the db when the main routine ends

	ginEngine := gin.Default()
	ginEngine.Use(middleware.Logger()) //Logger middleware

	dbhandler.ConnectDatabase()  // initialize db connection from the dbhandler package
	router.InitRoutes(ginEngine) //Initialize routes from the routes package

	ginEngine.Run() // Running on the default port 8080
}
