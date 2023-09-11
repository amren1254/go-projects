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

		if id, isValidUserCredentials, err := db.VerifyUserCredential(user); isValidUserCredentials {
			//generate and assign token
			token := model.TokenClaims{
				Id:       id,
				Username: user.Username,
			}
			generatedTokens, err := auth.GenerateAuthToken(token)
			if err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"SuccessMessage": "UserLoggedIn", "token": generatedTokens})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"WarningMessage": "Wrong Credential -" + err.Error()})
		}
	}
}

func GetProfile(db database.DatabaseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		//get all the user detail and account detail
		tokenClaims, err := auth.ExtractUsernameAndIdFromTokenClaims(c)
		if err != nil {
			log.Println("Error extracting username claim")
		}
		user, err := db.RetriveUser(tokenClaims)
		if err != nil {
			log.Println("error getting user details")
		}
		account, err := db.RetrieveAccountDetails(tokenClaims)
		if err != nil {
			log.Println("error getting account details")
		}
		var profile model.Profile
		profile.User = user
		profile.Account = account

		c.JSON(200, profile)
	}
}

func UpdateProfile(db database.DatabaseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
