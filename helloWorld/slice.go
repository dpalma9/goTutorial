package main

import "fmt"
// Los slices son como los arrays pero mas potentes.
// Pertimen realizar mas operaciones que los arrays.

func main() {

    // Para crear un slice vacio podemos usar make
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // Los slices se puede copiar tambien
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // el operador slice funciona como en python, para crear un array y rellenar entre dos valores
    l := s[2:5]
    fmt.Println("sl1:", l)

    // se puede tambien crear valores hasta pero excluyendo
    l = s[:5]
    fmt.Println("sl2:", l)
    // o valores hasta pero incluyendo
    l = s[2:]
    fmt.Println("sl3:", l)

    // Igual que el arary, los podemos crear multi-dimensionales
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
