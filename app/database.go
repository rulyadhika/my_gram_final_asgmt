package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	appConfig := GetAppConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", appConfig.DB_HOST, appConfig.DB_PORT, appConfig.DB_USERNAME, appConfig.DB_PASSWORD, appConfig.DB_DATABASE)

	db, err := sql.Open(appConfig.DB_DRIVER, dsn)

	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic(err)
	}

	return db
}
