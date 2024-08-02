package routes

import (
	"github.com/ahmed-afzal1/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	route := r.Group("users")
	{
		route.GET("/", controllers.GetAllUsers)
		route.POST("/", controllers.CreateUser)
		route.GET("/:id", controllers.FindUser)
		route.PUT("/:id", controllers.UpdateUser)
		route.DELETE("/:id", controllers.DeleteUser)
	}

	postRoute := r.Group("posts")
	{
		postRoute.GET("/", controllers.GetAllPost)
		postRoute.POST("/", controllers.CreatePost)
		postRoute.GET("/:id", controllers.FindPost)
		postRoute.PUT("/:id", controllers.UpdatePost)
		postRoute.DELETE("/:id", controllers.DeletePost)
	}
}
