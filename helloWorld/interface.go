package main

import (
    "fmt"
    "math"
)

// Una interfaz es una coleccion de metodos
// Siendo un metodo una coleccion de estructuras y funciones
type geometry interface {
    area() float64
    perim() float64
}

// cosas de rectangulos
type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// cosas de circulos
type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// usamos la interfaz
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    // es cada llamada, go sabe que interfaz usar
    measure(r)
    measure(c)
}
