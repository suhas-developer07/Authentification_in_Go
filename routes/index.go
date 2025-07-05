package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhas-developer07/Authentification_in_Go/routes/handlers"
)



func MountRoutes() *gin.Engine {
	handler := gin.Default();

	handler.POST("/signin",handlers.Signin)
	handler.POST("/signup",handlers.Signup)
	handler.GET("/",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"server working fine",
		})
	})

	return handler
}