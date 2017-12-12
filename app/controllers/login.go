package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

// Login ...
func Login(c *gin.Context) {
	OutputDataJSON(c, 200, "JWT Token", gin.H{"jwt": "jwt"})
	return
}
