package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() error {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Erro ao carregar arquivo .env: %v\n", err)
		return err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Printf("Error connection file .env: %v/n", err)
		return err
	}

	if err = DB.Ping(); err != nil {
		log.Printf("Error ping database:%v\n", err)
		return err
	}

	fmt.Printf("Connecting database in " + dbname + " success!")
	return nil
}
