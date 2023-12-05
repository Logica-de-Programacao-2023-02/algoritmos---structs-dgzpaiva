package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func (f *Funcionario) aumentarSalarioPorcentagem(quantidade float64) {
	f.salario *= (1 + quantidade/100)
}

func (f *Funcionario) diminuirSalarioPorcentagem(quantidade float64) {
	f.salario *= (1 - quantidade/100)
}

func (f *Funcionario) tempoDeServico() int {
	idadeInicioTrabalho := 18
	idadeAtual := f.idade
	return idadeAtual - idadeInicioTrabalho
}

func main() {
	funcionario := Funcionario{nome: "João", salario: 3000.0, idade: 30}

	fmt.Println("Salário antes do aumento:", funcionario.salario)
	funcionario.aumentarSalarioPorcentagem(10)
	fmt.Println("Salário após aumento de 10%:", funcionario.salario)

	fmt.Println("Salário antes da redução:", funcionario.salario)
	funcionario.diminuirSalarioPorcentagem(5)
	fmt.Println("Salário após redução de 5%:", funcionario.salario)

	fmt.Println("Tempo de serviço:", funcionario.tempoDeServico(), "anos")
}
