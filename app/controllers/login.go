package controllers

import (
	"log"

	"github.com/XanderDwyl/payrollspace.com/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// Login ...
func Login(c *gin.Context) {
	var LoginRequest models.LoginRequest

	if err := c.BindJSON(&LoginRequest); err != nil {
		log.Println(err)
		OutputJSON(c, 200, err.Error())
		return
	}

	user, err := LoginRequest.Login()
	if err != nil {
		OutputJSON(c, 200, err.Error())
		return
	}
	token, err := user.CreateJWToken()
	if err != nil {
		OutputJSON(c, 200, err.Error())
		return
	}
	OutputDataJSON(c, 200, "login ok", gin.H{
		"token":   token,
		"user_id": user.ID,
	})
	return
}
