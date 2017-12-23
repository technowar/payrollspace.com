package middleware

import (
	"log"
	"strings"

	"github.com/XanderDwyl/payrollspace.com/app/config"
	"github.com/XanderDwyl/payrollspace.com/app/controllers"
	"github.com/XanderDwyl/payrollspace.com/app/models"
	jwt "github.com/dgrijalva/jwt-go"

	"gopkg.in/gin-gonic/gin.v1"
)

// APIAuth ...
func APIAuth() gin.HandlerFunc {
	loginNotRequiredPaths := map[string]bool{
		"/check":      true,
		"/api/logout": true,
		"/api/signup": true,
		"/api/login":  true,
		"/api/hot":    true,
	}

	loginNotRequiredPathPrefixes := map[string]bool{
		"/api/verify": true,
		"/sitemap":    true,
	}

	return func(c *gin.Context) {
		if loginNotRequiredPaths[c.Request.URL.Path] {
			c.Next()
			return
		}

		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}

		for key, value := range loginNotRequiredPathPrefixes {
			if value {
				if strings.HasPrefix(c.Request.URL.Path, key) {
					c.Next()
					return
				}
			}
		}

		tokenString := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)

		if len(tokenString) == 0 {
			log.Println("Token String not found")

			c.Set("is_login", false)
			c.Next()
			return
		}

		user := models.JWTUser{}
		_, err := jwt.ParseWithClaims(tokenString, &user, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSalt()), nil
		})

		if err != nil {
			log.Println("JWT Error")

			c.Set("is_login", false)
			c.Next()
			return
		}

		if user.UserID == 0 || user.Username == "" {
			controllers.OutputErrorJSON(c, "User not found")
			c.Abort()
			return
		}

		c.Set("is_login", true)
		c.Set("access_token", user.AccessToken)
		c.Set("me", user)
		c.Next()
	}
}
