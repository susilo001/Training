package router

import (
	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/risk-profile-assement/handler"
	"github.com/susilo001/golang-advance/risk-profile-assement/middleware"
)

// SetupRouter initializes and sets up routes for the application
func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler, submissionHandler handler.ISubmissionHandler) {
	// Public endpoints for users
	usersPublicEndpoint := r.Group("/users")
	{
		usersPublicEndpoint.GET("/:id", userHandler.GetUser)
		usersPublicEndpoint.GET("", userHandler.GetAllUsers)
		usersPublicEndpoint.POST("", userHandler.CreateUser)
	}

	// Private endpoints for users with authentication middleware
	usersPrivateEndpoint := r.Group("/users")
	usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	{
		usersPrivateEndpoint.PUT("/:id", userHandler.UpdateUser)
		usersPrivateEndpoint.DELETE("/:id", userHandler.DeleteUser)
	}

	// Endpoints for submissions
	submissionsEndpoint := r.Group("/submissions")
	{
		submissionsEndpoint.POST("", submissionHandler.CreateSubmission)
		submissionsEndpoint.PUT("/:id", submissionHandler.UpdateSubmission)
		submissionsEndpoint.DELETE("/:id", submissionHandler.DeleteSubmission)
		submissionsEndpoint.GET("", submissionHandler.GetAllSubmissions)
		submissionsEndpoint.GET("/:id", submissionHandler.GetSubmissionByID)
	}
}
