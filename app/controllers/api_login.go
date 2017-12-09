package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func APILogin(c *gin.Context) {
	OutputDataJSON(c, 200, "JWT Token", gin.H{"jwt": "jwt"})
	return
}
