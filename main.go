package main

import (
	"log"
	"os"
	"time"

	"github.com/XanderDwyl/payrollspace.com/app/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	router := gin.Default()
	allowOrigins := []string{
		"http://payrollspace.com",
		"https://payrollspace.com",
	}

	if os.Getenv("MODE") != "production" {
		allowOrigins = append(allowOrigins, "http://localhost:3000")
		allowOrigins = append(allowOrigins, "http://localhost:3001")
		allowOrigins = append(allowOrigins, "http://127.0.0.1:3000")
		allowOrigins = append(allowOrigins, "http://127.0.0.1:3001")
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Accept", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	store := sessions.NewCookieStore([]byte("Lod5c5F"))

	router.Use(sessions.Sessions("mysession", store))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.Use(gin.Recovery())

	initializeRoutes(router)

	router.Run(":3000")
}

func initializeRoutes(origRouter *gin.Engine) {
	router := origRouter.Group("")

	router.GET("/", controllers.AppIndex)
}
