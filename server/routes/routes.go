package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/controllers"
	middleware "github.com/hyperyuri/webapi-with-go/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user") 
		{
			user.POST("/", controllers.CreateUser)
		}
		books := main.Group("books", middleware.Auth())
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		main.POST("login", controllers.Login)
	}

	return router
}
