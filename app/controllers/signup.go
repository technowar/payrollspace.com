package controllers

import (
	"github.com/XanderDwyl/payrollspace.com/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func Signup(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		OutputJSON(c, "error", err.Error())
		return
	}

	user, err = user.Create()
	if err != nil {
		OutputJSON(c, "error", err.Error())
		return
	}

	OutputJSON(c, "success", "User Created")
}
