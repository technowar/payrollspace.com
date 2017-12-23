package controllers

import "gopkg.in/gin-gonic/gin.v1"

// OutputJSON ...
func OutputJSON(c *gin.Context, status, msg string) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": msg,
	})
}

// OutputDataJSON ...
func OutputDataJSON(c *gin.Context, status, msg string, data gin.H) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": msg,
		"data":    data,
	})
}
