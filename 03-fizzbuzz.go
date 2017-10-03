package main

	import "fmt"

	func main() {
	fmt.Println("Welcome to fizzbuzz!")

	//create an array/slice assign it to variable c
	c := make([]int, 100)

		//loop over the array
		for i := range c {

			//increment the counter
			d := i + 1

			//calculate if variable d is divisible by 3
			threes := d%3 == 0
			//calculate if variable d is divisible by 5
			fives := d%5 == 0
			
			//compare int d against variables threes and fives
			//prints word according to condition met in if else statement
			if threes && fives {
				fmt.Println("FizzBuzz")
			} else if threes {
				fmt.Println("Fizz")
			} else if fives {
				fmt.Println("Buzz")
			} else {
				fmt.Println(d)
			}
		}
	}