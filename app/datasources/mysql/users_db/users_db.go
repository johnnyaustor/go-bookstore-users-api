package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlHost     = "mysql_users_host"
	mysqlSchema   = "mysql_users_schema"
	mysqlUsername = "mysql_users_username"
	mysqlPassword = "mysql_users_password"
)

var (
	Client *sql.DB

	host     = os.Getenv(mysqlHost)
	schema   = os.Getenv(mysqlSchema)
	username = os.Getenv(mysqlUsername)
	password = os.Getenv(mysqlPassword)
)

func init() {
	log.Println("try to connect database...")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err!=nil {
		panic(err)
	}

	if err = Client.Ping(); err!=nil {
		panic(err)
	}

	log.Println("Database Connected!")
}
