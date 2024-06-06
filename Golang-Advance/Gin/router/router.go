package router

import (
	"github.com/gin-gonic/gin"
	"github.com/susilo001/training-gin-framework/handler"
	"github.com/susilo001/training-gin-framework/middleware"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/:id", handler.GetUser)
	usersPublicEndpoint.GET("/", handler.GetAllUsers)

	usersPrivateEndpoint := r.Group("/users")
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	usersPrivateEndpoint.POST("/", handler.CreateUser)
	usersPrivateEndpoint.PUT("/:id", handler.UpdateUser)
	usersPrivateEndpoint.DELETE("/:id", handler.DeleteUser)

	// Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndpoint := r.Group("/private")
	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		privateEndpoint.POST("/post", handler.PostHandler)
	}
}
