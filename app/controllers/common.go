package controllers

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

const baseURL = "https://payrollspace.com"

func OutputErrorJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "error",
		"message": msg,
	})
}

func OutputSuccessJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "sucess",
		"message": msg,
	})
}

func OutputSuccessDataJSON(c *gin.Context, msg string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": msg,
		"data":    data,
	})
}

func OutputNotFound(c *gin.Context, tmpl string, data gin.H) {
	data["CFConnectingIP"] = c.Request.Header.Get("CF-Connecting-IP")
	data["XForwardedFor"] = c.Request.Header.Get("X-Forwarded-For")
	data["ClientIP"] = c.ClientIP()
	data["Host"] = c.Request.Host

	c.HTML(http.StatusNotFound, tmpl, data)
}

func OutputInternalServerError(c *gin.Context, tmpl string, data gin.H) {
	data["CFConnectingIP"] = c.Request.Header.Get("CF-Connecting-IP")
	data["XForwardedFor"] = c.Request.Header.Get("X-Forwarded-For")
	data["ClientIP"] = c.ClientIP()
	data["Host"] = c.Request.Host

	c.HTML(http.StatusInternalServerError, tmpl, data)
}
