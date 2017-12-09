package config

import (
	"os"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

const (
	envJWTSaltName    = "JWT_SALT"
	additionalJWTSalt = "tdpWjinLxsqU7FVZbVYkQmeuasfeasAeSDf"
)

var jwtSalt string
var envChecked bool

// init ...
func init() {
	CheckEnvs()
}

// CheckEnvs ...
func CheckEnvs() error {
	if envChecked {
		return nil
	}

	envJWTSalt := os.Getenv(envJWTSaltName)

	if envJWTSalt == "" {
		envJWTSalt = "asdfljasdfhjlkjkadsfjkajsdfjkhasdfkjakjsdhfhkjasdfk"
	}

	jwtSalt = envJWTSalt + additionalJWTSalt
	envChecked = true

	return nil
}

// GetCallbackURL ...
func GetCallbackURL(c *gin.Context) string {
	hostname := GetCallbackHostname(c)
	scheme := "http"

	if strings.Contains(strings.ToLower(c.Request.URL.Scheme), "https") {
		scheme = "https"
	}

	return scheme + "://" + hostname + "/callback"
}

// GetCallbackHostname ...
func GetCallbackHostname(c *gin.Context) string {
	return "localhost:3000"
}

// GetJWTSalt ...
func GetJWTSalt() string {
	if envChecked == false {
		os.Exit(1)
	}

	return jwtSalt
}
