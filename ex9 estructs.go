package main

import (
	"fmt"
)

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.notas = append(a.notas, nota)
}

func (a *Aluno) removerNota(index int) {
	if index >= 0 && index < len(a.notas) {
		a.notas = append(a.notas[:index], a.notas[index+1:]...)
	}
}

func (a Aluno) calcularMedia() float64 {
	total := 0.0
	for _, nota := range a.notas {
		total += nota
	}
	if len(a.notas) > 0 {
		return total / float64(len(a.notas))
	}
	return 0.0
}

func (a Aluno) imprimirInfo() {
	fmt.Printf("Nome: %s\nIdade: %d\n", a.nome, a.idade)
	media := a.calcularMedia()
	fmt.Printf("Média das notas: %.2f\n", media)
}

func main() {
	aluno := Aluno{
		nome:  "João",
		idade: 20,
		notas: []float64{8.5, 7.0, 9.5}, 
	}

	aluno.imprimirInfo()

	aluno.adicionarNota(6.0)
	aluno.imprimirInfo()

	aluno.removerNota(1)
	aluno.im
