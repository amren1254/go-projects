package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"netbanking/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// const (
// 	username = "postgres" //os.Getenv("DB_USERNAME")
// 	password = "root"     //os.Getenv("DB_PASSWORD")
// 	hostname = "127.0.0.1"
// 	port     = 5432
// 	dbname   = "netbanking"
// )

// func init() {
// 	//initializing database here
// 	NewDatabaseRepository(&sql.DB{}).InitDatabaseConnection()
// }

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
	// return newDb
}

func (db DatabaseRepository) InsertUser(user model.User) {
	id := uuid.New()
	query := `INSERT INTO "user"("id", "name", "username", "password", "status", "phone", "email", "created_at") VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	insert, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	//encrypt password before storing
	encryptedPassword, err := encrypt(user.Password)
	if err != nil {
		log.Println(err)
	}
	//get utc
	var datetime = time.Now().UTC()

	response, err := insert.Exec(
		id, user.Name,
		user.Username,
		encryptedPassword,
		user.Status,
		user.Phone,
		user.Email,
		datetime.Format(time.RFC3339),
	)
	insert.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println(response.RowsAffected())
}

func (db DatabaseRepository) retrieveUserDetails(user model.Login) (retrievedUserCredential model.Login, err error) {

	rows, err := db.DB.Query(`SELECT "username", "password" FROM "user" WHERE "username" = $1`, user.Username)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&retrievedUserCredential.Username, &retrievedUserCredential.Password)
		if err != nil {
			log.Println(err)
			return retrievedUserCredential, err
		}
	}
	return retrievedUserCredential, nil
}
