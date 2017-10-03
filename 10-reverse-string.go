package main

import "fmt"

func main() {
    s := "Hello, playground"
    fmt.Println(s)

    fmt.Println(Reverse(s))

}

func Reverse(s string) string {
    var reverse string
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}