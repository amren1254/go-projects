package database

import (
	"log"
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

func (db DatabaseRepository) retrieveUserLoginInfo(user model.Login) (id uuid.UUID, retrievedUserCredential model.Login, err error) {

	rows, err := db.DB.Query(`SELECT "id", "username", "password" FROM "user" WHERE "username" = $1`, user.Username)
	if err != nil {
		log.Println(err)
	}
	// var id uuid.UUID
	for rows.Next() {
		err = rows.Scan(&id, &retrievedUserCredential.Username, &retrievedUserCredential.Password)
		if err != nil {
			log.Println(err)
			return id, retrievedUserCredential, err
		}
	}
	return id, retrievedUserCredential, nil
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

func (db DatabaseRepository) RetriveUser(token model.TokenClaims) (retrievedUser model.User, err error) {
	rows, err := db.DB.Query(`SELECT "username","email","phone","status" FROM "user" WHERE "username" = $1 AND "id" = $2`, token.Username, token.Id)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		status := []uint8{}
		err = rows.Scan(&retrievedUser.Username, &retrievedUser.Email, &retrievedUser.Phone, &status)
		retrievedUser.Status = model.Status(string(status))
		if err != nil {
			log.Println(err)
			return retrievedUser, err
		}
	}
	return retrievedUser, nil
}

func (db DatabaseRepository) RetrieveAccountDetails(token model.TokenClaims) (retrievedAccount model.Account, err error) {
	rows, err := db.DB.Query(`SELECT "account_number", "account_type","total_amount" FROM "account" WHERE "id" = $1`, token.Id)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&retrievedAccount.Account_Number, &retrievedAccount.Account_Type, &retrievedAccount.Total_Amount)
		if err != nil {
			log.Println(err)
			return retrievedAccount, err
		}
	}
	return retrievedAccount, nil
}
