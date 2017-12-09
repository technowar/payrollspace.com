package controllers

import "gopkg.in/gin-gonic/gin.v1"

// AppIndex ...
func AppIndex(c *gin.Context) {
	OutputJSON(c, 200, "Welcome to payroll space!")
	return
}
