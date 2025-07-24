package controllers

import (
	"go-test/models"
	"go-test/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginUser := models.LoginRequest{}

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	if loginUser.Username == "" || loginUser.Password == "" {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Username and password must not be empty",
		})
		return
	}

	user, err := models.FindUser(loginUser.Username)
	if err != nil {
		log.Println("Error FindUser:", err.Error())
		c.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "User is not registered",
		})
		return
	}

	if !utils.VerifyPassword(user.Password, loginUser.Password) {
		c.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Invalid username or password",
		})
		return
	}
	token := utils.GenereateTokens(user.ID)

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Login.",
		Results: token,
	})

}
