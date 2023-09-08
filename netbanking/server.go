package main

import (
	"log"
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
	r := router.InitRoute()
	log.Fatal(r.Run(":8080"))
}
