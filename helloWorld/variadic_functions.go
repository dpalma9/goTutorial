package main

import "fmt"

// Funciones que pueden ser llamadas con cualquier numero de variables
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
