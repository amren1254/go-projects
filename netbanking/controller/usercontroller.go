package controller

import (
	"log"
	"net/http"

	"netbanking/auth"
	"netbanking/database"
	"netbanking/model"

	"github.com/gin-gonic/gin"
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
		db.InsertUser(user)
		c.IndentedJSON(http.StatusCreated, Response{Message: "User Created"})
	}
}

func Login(db database.DatabaseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {

		user := model.Login{}
		if err := c.BindJSON(&user); err != nil {
			log.Println(err)
		}

		isValidUserCredentials, err := db.VerifyUserCredential(user)
		if err != nil {
			log.Println(err)
		}
		if isValidUserCredentials {
			//generate and assign token
			token, err := auth.GenerateAuthToken(user.Username)
			if err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"SuccessMessage": "UserLoggedIn", "token": token})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"WarningMessage": "Wrong Credential"})
		}
	}

}
