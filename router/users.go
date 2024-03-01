package router

import (
	"project3/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(ginEngine *gin.Engine) {
	ginEngine.GET("/status/", handlers.Status)

	users := ginEngine.Group("/users")
	{
		users.POST("/", handlers.CreateUser)
		users.GET("/", handlers.GetAllUsers)
		users.GET("/:user-id/", handlers.GetIndividualUser)
		users.PUT("/:user-id/", handlers.UpdateUser)
		users.DELETE("/:user-id/", handlers.DeleteUser)
	}
}
