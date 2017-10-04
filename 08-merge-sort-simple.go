package main

import (
	"fmt"
)

func main() {
	sliceOne := []int{1, 2, 3, 4, 5, 6} // sorted
	sliceTwo := []int{7, 8, 9, 10, 11, 12} // sorted
	sliceThree := []int{20, 21, 22, 23, 24, 25} // sorted
	sliceOne = append(sliceOne, sliceTwo...)
	sliceOne = append(sliceOne, sliceThree...)
	fmt.Println(sliceOne)
}