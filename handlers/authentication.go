package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaswanthsaivendra/prod_mang/helper"
	"github.com/yaswanthsaivendra/prod_mang/model"
)

func Register(c *gin.Context) {
	var input model.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username:  input.Username,
		Password:  input.Password,
		Mobile:    input.Mobile,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	}

	savedUser, err := user.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(c *gin.Context) {
	var input model.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
