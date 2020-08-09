package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var mode int
	mode = 100
	fmt.Println("What range of numbers would you like to guess from?")
	fmt.Scan(&mode)
	//
	fmt.Println("Guess the number 1 out of", mode)
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(mode)

	var guess int
	var guessCount int
	guessCount = 1

	for {
		fmt.Println("enter your guess :  ")
		fmt.Scan(&guess)
		if guess > mode {
			fmt.Printf("Your guess of %d was greater than the maximum number of %d \n", guess, mode)

		} else if guess < secretNumber {
			fmt.Println("The number is higher")
			guessCount++
		} else if guess > secretNumber {
			fmt.Println("The number is lower")
			guessCount++
		} else {
			fmt.Printf("You won with %v guesses \n", guessCount)
			break
		}
	}
}
