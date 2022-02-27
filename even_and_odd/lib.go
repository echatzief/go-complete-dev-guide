package main

import (
    "fmt"
)

func isOddOrEven(num int){
    remainder := num % 2
    if remainder == 0 {
        fmt.Println(num," : Even")
    } else {
        fmt.Println(num," : Odd")
    }
}
