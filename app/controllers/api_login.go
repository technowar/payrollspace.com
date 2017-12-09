package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func APILogin(c *gin.Context) {
	OutputSuccessDataJSON(c, "JWT Token", gin.H{"jwt": "jwt"})
	return
}
