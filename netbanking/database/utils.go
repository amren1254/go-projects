package database

import (
	"math/rand"
	"strconv"
)

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
