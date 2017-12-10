package controllers

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func OutputJSON(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"status":  http.StatusText(code),
		"code":    code,
		"message": msg,
	})
}

func OutputDataJSON(c *gin.Context, code int, msg string, data gin.H) {
	c.JSON(code, gin.H{
		"status":  http.StatusText(code),
		"code":    http.StatusOK,
		"message": msg,
		"data":    data,
	})
}
