package models

import (
	"fmt"
	"log"
	"os"
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
	if dbPass != "" {
		dbPass = ":" + dbPass
	}

	tryOnlyOnce := getenvWithDefault("TRY_ONLY_ONCE", "")
	skipMigration := getenvWithDefault("SKIP_MIGRATION", "")
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

	var dberr error
	//var db *gorm.DB
	loop := 0
	for true {
		db, dberr = gorm.Open("mysql", dsn1)
		if dberr == nil {
			log.Printf("db connect OK: %s", dsn1)
			isConnected = true
			break
		}

		log.Print("DB Connection Error: dsn1=" + dsn1)

		log.Print(dberr.Error())
		if tryOnlyOnce != "" {
			return
		}

		time.Sleep(time.Millisecond * 3000)

		db, dberr = gorm.Open("mysql", dsn2)
		if dberr == nil {
			log.Printf("db connect OK: %s", dsn2)
			log.Printf("create database %s", dbName)
			db.Exec("CREATE DATABASE IF NOT EXISTS `" + dbName + "` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
			// might be error...?
			log.Printf("use dbname")
			db.Exec("use `" + dbName + "` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
			db.Close()
		} else {
			log.Print("DB Connection Error: dsn2=" + dsn2)
			log.Print(dberr.Error())
			loop++
			if loop > 300 {
				//give up to connect... even though no connect db.
				break
			}
		}
	}

	log.Print("FLUSH HOSTS;")
	err := db.Debug().Exec("FLUSH HOSTS;").Error
	if err != nil {
		log.Print(err)
	} else {
		log.Print("flush hosts => ok")
	}

	//ResetTables()
	if skipMigration != "" {
		log.Print("skip migration")
		//MigrateTables()
	} else {
		log.Printf("before auto migrate.")
		MigrateTables()
		log.Printf("after auto migrate..")
	}

}

// MigrateTables ...
func MigrateTables() {
	db.AutoMigrate(&User{})
}

// getenvWithDefault ...
func getenvWithDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
