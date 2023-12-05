package main

import "fmt"

type Triangulo struct {
    base   float64
    altura float64
}

func calcularAreaTriangulo(t Triangulo) float64 {
    return (t.base * t.altura) / 2
}

func main() {
    meuTriangulo := Triangulo{base: 5.0, altura: 8.0}
    area := calcularAreaTriangulo(meuTriangulo)
    fmt.Println("A área do triângulo é:", area)
}
