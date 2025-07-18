package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/reawphongsaphak/slOwOrk-Backend/errors"
	"github.com/reawphongsaphak/slOwOrk-Backend/models"
	"github.com/reawphongsaphak/slOwOrk-Backend/services"
)

func LoginUser(c *gin.Context) {
	var loginUser models.LoginUser

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(errors.ErrBadRequest.Code, gin.H{
			"error": errors.ErrBadRequest.Message,
		})
		return
	}

	tokenString, err := services.LoginUser(loginUser)
	if err != nil {
		c.JSON(errors.ErrInternalServer.Code, gin.H{
			"error": err,
		})
		return
	}

	//* set time out to 1 hour
	c.SetCookie("Authorization", tokenString, 60*60, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}