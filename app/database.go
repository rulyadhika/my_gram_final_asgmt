package app

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDB() *sql.DB {
	appConfig := GetAppConfig()

	dsn := fmt.Sprintf("host:%s port:%s database:%s username:%s password:%s", appConfig.DB_HOST, appConfig.DB_PORT, appConfig.DB_DATABASE, appConfig.DB_USERNAME, appConfig.DB_PASSWORD)

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
