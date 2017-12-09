package controllers

import "gopkg.in/gin-gonic/gin.v1"

func AppIndex(c *gin.Context) {
	OutputJSON(c, 200, "Welcome")
	return
}
