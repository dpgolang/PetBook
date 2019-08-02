package driver

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found:%v", err)
	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	var err error
	var db *sql.DB
	host := os.Getenv("SERVICE_HOST")
	port := os.Getenv("SERVICE_PORT_BD")
	user := os.Getenv("SERVICE_USER")
	password := os.Getenv("SERVICE_PASSWORD")
	dbname := os.Getenv("SERVICE_DBNAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	logFatal(err)
	err = db.Ping()
	if err != nil {
		logFatal(fmt.Errorf("can't ping, err: %s", err.Error()))
	}
	return db
}
