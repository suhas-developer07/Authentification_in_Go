package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhas-developer07/Authentification_in_Go/routes/handlers"
)



func MountRoutes() *gin.Engine {
	handler := gin.Default();

	handler.POST("/signin",handlers.Signin)
	handler.POST("/signup",handlers.Signup)

	return handler
}