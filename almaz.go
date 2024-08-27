package main

import (
        "fmt")
func main() {
        romeToArab := map[string]int{

                "I": 1,
                "II": 2,
                "III": 3,
        }
        ArabToRome := map[int]string{
                1: "I",
                2: "II",
                3: "III",
        }
        res := romeToArab["I"] + romeToArab["II"]
        fmt.Println(ArabToRome[res])
}
