package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammad-quanit/auth/models"
	"github.com/mohammad-quanit/auth/utils"
)

func Signup(c *gin.Context) {

	// Recieve User Information
	var user models.User

	if err := c.ShouldBindJSON(&user); err != err {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Chec if email already exists before creating new
	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	models.DB.Create(&user)
	c.JSON(200, gin.H{"success": "user created"})
}
