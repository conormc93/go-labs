package main

import (
    "fmt"
    "strings"
)

func main() {
    var input string
    fmt.Println("Enter string:")
    fmt.Scanf("%s\n", &input)
    input = strings.ToLower(input)
    fmt.Println(isPalindrome(input))
}

func isPalindrome(s string) string {
 mid := len(s) / 2
 last := len(s) - 1
 for i := 0; i < mid; i++ {
  if s[i] != s[last-i] {
   return "NO. It's not a Palimdrome."
  }
 }
 return "YES! You've entered a Palindrome"
}