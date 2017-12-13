package models

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var isConnected bool

// Setup ..
func Setup() {
	if isConnected == true {
		return
	}

	dbHost := getenvWithDefault("DB_HOST", "localhost")
	dbName := getenvWithDefault("DB_NAME", "payrollspace")
	dbUser := getenvWithDefault("DB_USER", "root")
	dbPass := getenvWithDefault("DB_PASS", "")
	dbPort := getenvWithDefault("DB_PORT", "3306")
	tryOnlyOnce := getenvWithDefault("TRY_ONLY_ONCE", "")
	skipMigration := getenvWithDefault("SKIP_MIGRATION", "")

	if dbPass != "" {
		dbPass = ":" + dbPass
	}

	dsn1 := fmt.Sprintf(
		"%s%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	dsn2 := fmt.Sprintf(
		"%s%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		"mysql",
	)

	var err error
	loop := 0

	for true {
		db, err = gorm.Open("mysql", dsn1)

		if err == nil {
			log.Printf("DB Connection Success: %s", dsn1)

			isConnected = true

			break
		}

		log.Printf("DB %s Connection Error: %s", dsn1, err.Error())

		if tryOnlyOnce != "" {
			return
		}

		time.Sleep(time.Millisecond * 3000)

		db, err = gorm.Open("mysql", dsn2)

		if err == nil {
			log.Printf("DB Connection Success: %s", dsn2)
			log.Printf("DB Creation:  %s", dbName)

			db.Exec("CREATE DATABASE IF NOT EXISTS `" + dbName + "` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
			db.Exec("use `" + dbName + "` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")

			db.Close()
		} else {
			log.Printf("DB %s Connection Error: %s", dsn2, err.Error())

			loop++

			if loop > 300 {
				break
			}
		}
	}

	err = db.Debug().Exec("FLUSH HOSTS;").Error

	if err != nil {
		log.Printf("Flush Hosts Error: %s", err)
	} else {
		log.Print("Flush Hosts Success")
	}

	//ResetTables()
	if skipMigration == "" {
		log.Printf("Migrating Tables")

		MigrateTables()
	}

}

// MigrateTables ...
func MigrateTables() {
	// users table
	db.AutoMigrate(&User{})

	// role_access table
	db.AutoMigrate(&RoleAccess{})

	// subscribers table
	db.AutoMigrate(&Subscribers{})
}

// getenvWithDefault ...
func getenvWithDefault(key string, def string) string {
	v := os.Getenv(key)

	if v == "" {
		return def
	}

	return v
}

// hashedPassword ...
func hashedPassword(rawPassword string) string {
	s := sha256.New()
	s.Write([]byte(rawPassword))
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}

// sanitizeEmail ...
func sanitizeEmail(email string) (string, error) {
	email = strings.ToLower(email)
	email = strings.Trim(email, "\r\n\t @!#$%^&*(){}[];'.,/'")
	if emailRegex.Match([]byte(email)) {
		return email, nil
	}
	return email, fmt.Errorf("invalid email")

}
