package controller

import (
	"log"
	"net/http"

	"netbanking/auth"
	"netbanking/database"
	"netbanking/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Response struct {
	Message string `json:"message"`
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"Ping": "Pong"})
}

func Signup(db database.DatabaseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			return
		}
		id := uuid.New()
		if isUserCreated := db.InsertUser(user, id); isUserCreated {
			log.Println("user created")
			//if user is created successfully, then create account
			if isAccountCreated := db.CreateAccount(id); isAccountCreated {
				log.Println("account created successfully")
			}
			c.IndentedJSON(http.StatusCreated, Response{Message: "User Created"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, Response{Message: "Unable to create user"})
		}
	}
}

func Login(db database.DatabaseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {

		user := model.Login{}
		if err := c.BindJSON(&user); err != nil {
			log.Println(err)
		}

		if isValidUserCredentials, err := db.VerifyUserCredential(user); isValidUserCredentials {
			//generate and assign token
			token, err := auth.GenerateAuthToken(user.Username)
			if err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"SuccessMessage": "UserLoggedIn", "token": token})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"WarningMessage": "Wrong Credential -" + err.Error()})
		}
	}
}

//for manual account creation
// func CreateAccount(db database.DatabaseRepository) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		account := model.Account{}
// 		if err := c.BindJSON(&account); err != nil {
// 			log.Println(err)
// 		}
// 		isAccountCreated := db.CreateAccount(account)
// 		if isAccountCreated {

// 		}
// 	}
// }

func GetProfile(db database.DatabaseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		//get all the user detail and account detail

	}
}

func UpdateProfile(db database.DatabaseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
