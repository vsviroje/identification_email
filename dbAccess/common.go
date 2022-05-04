package dbAccess

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB ...
var DB *sql.DB

// ConnectDB ...
func ConnectDB() error {
	db, err := sql.Open("mysql", "vsviroje:vvCO2142*#@tcp(127.0.0.1:3306)/training")
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	return err
}

func DisconnectDB() {
	DB.Close()
}
