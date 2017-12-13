package controllers

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// OutputJSON ...
func OutputJSON(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"status":  http.StatusText(code),
		"message": msg,
	})
}

// OutputDataJSON ...
func OutputDataJSON(c *gin.Context, code int, msg string, data gin.H) {
	c.JSON(code, gin.H{
		"status":  http.StatusText(code),
		"message": msg,
		"data":    data,
	})
}
