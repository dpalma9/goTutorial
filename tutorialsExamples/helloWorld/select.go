package main

import (
    "fmt"
    "time"
)

// Los selects se combinan con los channels y las goroutines
func main() {

    // Creamos dos channels que luego seleccionaremos
    c1 := make(chan string)
    c2 := make(chan string)

    // En estas dos goroutines, cada canal recibira un valor despues de un tiempo
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // Con el select lo que hacemos es esperar a los dos canales de forma simultanea
    // printando cada uno de ellos cuando llegue
    // Espera por el for, porque realmente el select cogeria el valor de la primera rutina en terminar
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
