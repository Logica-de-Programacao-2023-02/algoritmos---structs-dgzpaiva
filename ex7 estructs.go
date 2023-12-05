package main

import (
	"fmt"
)

type Animal struct {
	nome   string
	especie string
	idade  int
	som    string
}

func (a *Animal) alterarSom(novoSom string) {
	a.som = novoSom
}

func (a Animal) imprimirInfoSom() {
	fmt.Printf("Nome: %s\nEspécie: %s\nIdade: %d\nSom que faz: %s\n", a.nome, a.especie, a.idade, a.som)
}

func main() {
	animal := Animal{nome: "Rex", especie: "Cachorro", idade: 5, som: "Latido"}

	animal.imprimirInfoSom()

	fmt.Println("Alterando o som...")
	animal.alterarSom("Miado")

	animal.imprimirInfoSom()
}
