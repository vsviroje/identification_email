package dbAccess

import (
	"database/sql"
	"identification_email/utility/logger"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB ...
var DB *sql.DB

// ConnectDB ...
func ConnectDB() error {
	logger.I("ConnectDB invoked")
	defer logger.I("ConnectDB returned")

	db, err := sql.Open("mysql", "vsviroje:vvCO2142*#@tcp(127.0.0.1:3306)/training")
	if err != nil {
		logger.E("sql.Open failed", err)
		log.Fatal(err)
	}

	DB = db

	return err
}

func DisconnectDB() {
	logger.I("DisconnectDB invoked")
	defer logger.I("DisconnectDB returned")

	DB.Close()
}
