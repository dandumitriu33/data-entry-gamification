package receipts_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client   *sql.DB
	username = os.Getenv("MYSQL_DEV_USERNAME")
	password = os.Getenv("MYSQL_DEV_PASSWORD")
	host     = "127.0.0.1:3306"
	schema   = "vehicleregistration"
)

func init() {

	// username:password@tcp(host)/user_schema
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}