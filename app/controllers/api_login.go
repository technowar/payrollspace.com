package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

// // APILoginRequest ...
// type APILoginRequest struct {
// 	AccessToken string `json:"access_token"`
// }

// APILogin ...
func APILogin(c *gin.Context) {
	OutputOKDataJSON(c, "sample msg", gin.H{"jwt": "jwt"})
	return
}
