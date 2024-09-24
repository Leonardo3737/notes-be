package routes

import (
	userController "github.com/Leonardo3737/notes-be/src/controller/User"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {
	// USER
	router.GET("/user/:id", userController.FindUserById)
	router.GET("/user/getByEmail/:email", userController.FindUserByEmail)
	router.POST("/user/", userController.CreateUser)
	router.PUT("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)
}
