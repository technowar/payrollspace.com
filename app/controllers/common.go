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

func OutputError(c *gin.Context, code int, msg string, data gin.H) {
	data["CFConnectingIP"] = c.Request.Header.Get("CF-Connecting-IP")
	data["XForwardedFor"] = c.Request.Header.Get("X-Forwarded-For")
	data["ClientIP"] = c.ClientIP()
	data["Host"] = c.Request.Host

	c.HTML(code, msg, data)
}
