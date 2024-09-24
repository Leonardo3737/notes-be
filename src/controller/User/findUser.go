package userController

import (
	rest_err "github.com/Leonardo3737/notes-be/src/config/exception"
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("test Error")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {}
