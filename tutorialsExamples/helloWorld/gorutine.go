package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
// Las goroutines son hilos ligeros y que se invocan con la llamada go function()
// Ej: go calc(num)
func main() {

    f("direct")

    go f("goroutine")

    //go func(msg string) {
    //    fmt.Println(msg)
    //}("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}
