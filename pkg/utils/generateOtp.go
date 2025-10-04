package utils

import (
	"math/rand"
	"strconv"
)

func GenerateOtp() string {
	// Generate a random 6-digit number (between 100000 and 999999)
	randomNumber := rand.Intn(900000) + 100000
	return strconv.Itoa(randomNumber)
}
