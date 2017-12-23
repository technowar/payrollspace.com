package controllers

import (
	"github.com/XanderDwyl/payrollspace.com/app/models"

	"gopkg.in/gin-gonic/gin.v1"
)

// Signup ...
func Signup(c *gin.Context) {
	var err error
	var user models.User

	if err = c.BindJSON(&user); err != nil {
		OutputJSON(c, "error", err.Error())
		return
	}

	user, err = user.Create()

	if err != nil {
		OutputJSON(c, "error", err.Error())
		return
	}

	// token, err := user.CreateJWToken()
	// if err != nil {
	// 	log.Println(err)
	// 	OutputJSON(c, 200, err.Error())
	// 	return
	// }

	// TODO: send email notification

	OutputDataJSON(c, "success", "login ok", gin.H{
		"token":     "token",
		"user_uuid": "user.UUID",
	})
	return
}
