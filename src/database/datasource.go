package database

import (
	"database/sql"
	"fmt"
	"gofiber-sqlc/src/pkg/config"
	"gofiber-sqlc/src/pkg/utils"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type FilteredQuery struct {
	Page     string
	PageSize string
}

var DB *sql.DB

var (
	host     = config.Env("DB_HOST", "127.0.01")
	port, _  = strconv.Atoi(config.Env("DB_PORT", "5432"))
	dbname   = config.Env("DB_DATABASE", "db_example")
	username = config.Env("DB_USERNAME", "postgres")
	password = config.Env("DB_PASSWORD", "postgres")
)

func ConnectDB() {
	// Conection to the database
	var connect string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db
	msg := fmt.Sprintf("Connection %s has been established successfully.", dbname)

	logMessage := utils.PrintLog("Database", msg)
	fmt.Println(logMessage)
}
