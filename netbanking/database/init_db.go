package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type DatabaseRepository struct {
	DB *sql.DB
}

func NewDatabaseRepository(database *sql.DB) *DatabaseRepository {
	return &DatabaseRepository{
		DB: database,
	}
}

func dataSourceName(databaseName string) string {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOSTNAME"), port, os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}

func (db *DatabaseRepository) InitDatabaseConnection() {
	var err error
	dbname := os.Getenv("DB_NAME")
	newDb, err := sql.Open("postgres", dataSourceName(dbname))
	if err != nil {
		log.Printf("Error while opening connection with database")
		return
	}
	//defer DB.Close()

	newDb.SetMaxOpenConns(10)
	newDb.SetMaxIdleConns(10)
	newDb.SetConnMaxLifetime(time.Minute * 5)
	db.DB = newDb
}
