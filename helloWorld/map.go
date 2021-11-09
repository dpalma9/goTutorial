package main

import "fmt"

// un mapa es parecido a un diccionario de Python
func main() {

    // un mapa no deja de ser variable con clave-valor.
    m := make(map[string]int)
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m)

    // Podemos recuperar valores de una clave
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // Y borrar clave/valor
    delete(m, "k2")
    fmt.Println("map:", m)

    // Esto seria un enemplo de poner un valor por defecto en caso de no encontrar una clave
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
}
