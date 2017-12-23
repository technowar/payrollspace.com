package controllers

import (
	"github.com/XanderDwyl/payrollspace.com/app/models"

	"gopkg.in/gin-gonic/gin.v1"
)

// Login ...
func Login(c *gin.Context) {
	var LoginRequest models.LoginRequest

	if err := c.BindJSON(&LoginRequest); err != nil {
		OutputJSON(c, "error", err.Error())
		return
	}

	user, err := LoginRequest.Login()

	if err != nil {
		OutputJSON(c, "error", err.Error())
		return
	}

	token, err := user.CreateJWToken()

	if err != nil {
		OutputJSON(c, "error", err.Error())
		return
	}

	OutputDataJSON(c, "success", "login ok", gin.H{
		"token":   token,
		"user_id": user.ID,
	})
	return
}
