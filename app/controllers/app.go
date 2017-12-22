package controllers

import "gopkg.in/gin-gonic/gin.v1"

// Index ...
func Index(c *gin.Context) {
	OutputJSON(c, "success", "Welcome to payroll space!")
	return
}
