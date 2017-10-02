package main

import "fmt"
import "math/rand"
import "time"

func main() {
	fmt.Print("Welcome to the Guessing Game\n")
	var guessCount int
	rand.Seed(time.Now().UTC().UnixNano())
	var answer = rand.Intn(100)
	fmt.Print("Enter a number between 1 & 100.\n")

    fmt.Println("Enter a number: ")
    var guess int
	fmt.Scanln(&guess)
	for guess != answer{
		if guess > answer{
			fmt.Println("Incorrect. Lower.")
			guessCount++
			fmt.Scanln(&guess)
		} else if guess < answer{
			fmt.Println("Incorrect. Higher.")
			guessCount++
			fmt.Scanln(&guess)
		} else if guess == answer{
			return
		}
	}
	fmt.Println("Correct!!")
	fmt.Printf("Number of guesses: %d", guessCount)

}