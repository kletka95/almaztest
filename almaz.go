package main

import "fmt"

// this is a comment

func main() {
    fmt.Println("Hello World:")
    var first int
    var sec int
    fmt.Scan(&first, &sec)
    fmt.Println(first+sec)
}
