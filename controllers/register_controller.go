package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/reawphongsaphak/slOwOrk-Backend/models"
	"github.com/reawphongsaphak/slOwOrk-Backend/errors"
	"github.com/reawphongsaphak/slOwOrk-Backend/services"

)

//* register new user
func RegisterNewUser(c *gin.Context) {
	var newUser models.User
	
	//* binding into data model
	if err := c.ShouldBindJSON(&newUser); err != nil {
		//* repond error if fail
		c.JSON(errors.ErrBadRequest.Code, gin.H{
			"error": errors.ErrBadRequest.Message,
		})
	}

	//* register new user
	result, err := services.RegisterNewUser(newUser)
	if err != nil {
		//* response error if fail
		c.JSON(errors.ErrInternalServer.Code, gin.H{"error": err.Error()})
		return
	}

	//* response message and inserted id
	c.JSON(errors.ErrCreated.Code, gin.H{
		"message": errors.ErrCreated.Message,
		"id":      result.InsertedID,
	})

	
}

