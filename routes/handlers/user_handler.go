package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "github.com/suhas-developer07/Authentification_in_Go/db/user_db"
	"github.com/suhas-developer07/Authentification_in_Go/models"
)


func Signin(ctx *gin.Context) {
   var Payload models.SigninPayload

   if err := ctx.ShouldBindJSON(&Payload) ; err != nil{
      ctx.JSON(http.StatusBadRequest, gin.H{
		"error" :true,
		"message":"unable to read the body",
	  })
   }

   user,err := userdb.UserRepository.CheckUserExist(Payload)

   if err!=nil {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error":true,
		"message" : err.Error(),
	})
	return
   }

   if Payload.Password != user.Password {
	ctx.JSON(http.StatusBadRequest,gin.H{
		"error":true,
		"message":"Invalid credentials",
	})
	return
   } 

   jwt := "123456"

   ctx.JSON(http.StatusOK,gin.H{
	"error":false,
	"message":"Login succesfull",
	"jwt":jwt,
   })
}

func Signup(ctx *gin.Context){
     var Payload models.SignupPayload

	 if err := ctx.ShouldBindJSON(&Payload); err!= nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" :true,
		    "message":"unable to read the body",
		})
		return
	}

	err := userdb.UserRepository.CreateUserQuery(Payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error" :true,
			"message" :err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"error" :false,
		"message":"user created successfully",
	})
	
}