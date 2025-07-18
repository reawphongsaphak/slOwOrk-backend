package routes

import (
	"github.com/reawphongsaphak/slOwOrk-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1") 
	{
		api.POST("/register", controllers.RegisterNewUser)
		api.POST("/login", controllers.LoginUser)
	}	
}