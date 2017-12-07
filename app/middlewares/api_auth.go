package middleware

import (
	"log"
	"strings"

	"github.com/XanderDwyl/payrollspace.com/app/config"
	"github.com/XanderDwyl/payrollspace.com/app/controllers"
	"github.com/XanderDwyl/payrollspace.com/app/models"
	jwt "github.com/dgrijalva/jwt-go"
	//"github.com/gin-gonic/gin"
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

		for path, b := range loginNotRequiredPathPrefixes {
			if b {
				if strings.HasPrefix(c.Request.URL.Path, path) {
					c.Next()
					return
				}
			}
		}

		tokenString := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
		if len(tokenString) == 0 {
			//apiv1handlers.OutputErrorJSON(c, "invalid sig")
			//c.Abort()
			log.Print("invalid sig")
			c.Set("is_login", false)
			c.Next()
			return
		}

		user := models.JWTUser{}
		_, err := jwt.ParseWithClaims(tokenString, &user, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSalt()), nil
		})

		if err != nil {
			//apiv1handlfers.OutputErrorJSON(c, "JWT error")
			//c.Abort()
			log.Print("JWT error")
			c.Set("is_login", false)
			c.Next()
			return
		}

		if user.UserID == 0 || user.Username == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{"err": "invalid sig or no such user"})
			controllers.OutputErrorJSON(c, "invalid sig or no such user")
			c.Abort()
			return
		}
		c.Set("is_login", true)
		c.Set("access_token", user.AccessToken)
		c.Set("me", user)
		c.Next()
	}

}
