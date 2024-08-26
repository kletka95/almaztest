package main

import "fmt"

// this is a comment

func main() {
    fmt.Println("Сюда ты можешь внести свои циферки и какие-нибудь знаки:")
    var first int
    var sec int
    
    fmt.Scan(&first, &sec)
    fmt.Println(first + sec)
    fmt.Println(first * sec)
}
