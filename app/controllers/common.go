package controllers

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

const (
	baseURL = "https://payrollspace.com"
)

// OutputErrorJSON ...
func OutputErrorJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "error",
		"message": msg,
	})
}

// OutputOKJSON ...
func OutputOKJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": msg,
	})
}

// OutputOKDataJSON ...
func OutputOKDataJSON(c *gin.Context, msg string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": msg,
		"data":    data,
	})
}

// OutputNotFound ...
func OutputNotFound(c *gin.Context, tmpl string, data gin.H) {
	data["CFConnectingIP"] = c.Request.Header.Get("CF-Connecting-IP")
	data["XForwardedFor"] = c.Request.Header.Get("X-Forwarded-For")
	data["ClientIP"] = c.ClientIP()
	data["Host"] = c.Request.Host
	c.HTML(http.StatusNotFound, tmpl, data)
}

// OutputInternalServerError ...
func OutputInternalServerError(c *gin.Context, tmpl string, data gin.H) {
	data["CFConnectingIP"] = c.Request.Header.Get("CF-Connecting-IP")
	data["XForwardedFor"] = c.Request.Header.Get("X-Forwarded-For")
	data["ClientIP"] = c.ClientIP()
	data["Host"] = c.Request.Host
	c.HTML(http.StatusInternalServerError, tmpl, data)
}
