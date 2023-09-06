package main

import (
	"database/sql"
	"fmt"

	"log"
	"netbanking/database"
	"netbanking/router"

	"github.com/joho/godotenv"
)

func init() {
	//initialize enviroment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading environment Vars - %v \n", err)
	}

}

func main() {
	fmt.Println(database.DatabaseRepository{DB: &sql.DB{}})
	r := router.InitRoute()
	r.Run(":8080")
}
