package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// NewRandomNumber returns a random number between 0 and "high".
func NewRandomNumber(high int) int {
	var number int
	for number == 0 {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		number = random.Intn(high)
	}
	return number
}

// GetUserGuessInt gets user's guess as an int.
func GetUserGuessInt() int {
	i, _ := strconv.ParseInt(GetUserGuess(), 10, 32)
	return int(i)
}

// GetUserGuess gets a user's guess.
func GetUserGuess() string {
	var guessString string
	fmt.Scanln(&guessString)
	return guessString
}
