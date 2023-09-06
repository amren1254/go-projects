package database

import (
	"errors"
	"fmt"
	"log"

	"netbanking/model"

	"golang.org/x/crypto/bcrypt"
)

func encrypt(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to encrypt password, please provide another password")
	}
	fmt.Println("hash to store ", string(hash))
	//store this hash into database
	return hash, nil
}

func (db DatabaseRepository) VerifyUserCredential(user model.Login) (bool, error) {
	// retrive password hash from database
	retrievedUser, err := db.retrieveUserDetails(user)
	if err != nil {
		log.Println(err)
	}

	//compare hash from database
	if err := bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(user.Password)); err != nil {
		return false, err
	}
	return true, nil
}
