package database

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"netbanking/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

//psql -U username -d myDataBase -a -f myInsertFile

func (db DatabaseRepository) InsertUser(user model.User, id uuid.UUID) bool {

	query := `INSERT INTO "user"("id", "name", "username", "password", "status", "phone", "email", "created_at") VALUES($1,$2,$3,$4,$5::text::user_status,$6,$7,$8)`
	insert, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return false
	}
	//encrypt password before storing
	encryptedPassword, err := encrypt(user.Password)
	if err != nil {
		log.Println(err)
		return false
	}
	//get utc
	var datetime = time.Now().UTC()

	response, err := insert.Exec(
		id,
		user.Name,
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
	if userCreate, err := response.RowsAffected(); userCreate == 1 {
		return true
	} else {
		log.Println(err)
		return false
	}
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

// generate account number
func generateAccountNumber() string {
	// Seed the random number generator with the current time.
	// rand.Seed(time.Now().UnixNano())

	// Generate a random 12 digit number.
	randomNumber := rand.Intn(1000000000000) + 1
	return strconv.Itoa(randomNumber)
	// Print the random number to the console.
	// fmt.Println(randomNumber)
}

func (db DatabaseRepository) CreateAccount(id uuid.UUID) bool {

	query := `INSERT INTO "account"
	("id", "account_number", "account_type", "total_amount", "created_at") 
	VALUES($1,$2,$3::text::account_type,$4,$5)`
	insert, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return false
	}
	//get utc
	var datetime = time.Now().UTC()

	response, err := insert.Exec(
		id,
		generateAccountNumber(),
		"savings",
		0,
		datetime.Format(time.RFC3339),
	)
	insert.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	if accountCreate, err := response.RowsAffected(); accountCreate == 1 {
		return true
	} else {
		log.Println(err)
		return false
	}
}
