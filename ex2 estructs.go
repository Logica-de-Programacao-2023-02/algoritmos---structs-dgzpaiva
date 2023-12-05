package main

import (
    "fmt"
)

type Endereco struct {
    rua     string
    numero  int
    cidade  string
    estado  string
}

type Pessoa struct {
    nome     string
    idade    int
    endereco Endereco
}

func imprimirEnderecoCompleto(p Pessoa) {
    fmt.Printf("Endereço de %s:\n", p.nome)
    fmt.Printf("Rua: %s, Número: %d\nCidade: %s, Estado: %s\n",
        p.endereco.rua, p.endereco.numero, p.endereco.cidade, p.endereco.estado)
}

func main() {
    enderecoExemplo := Endereco{
        rua:    "Rua das Flores",
        numero: 123,
        cidade: "Cidade Exemplo",
        estado: "Estado Exemplo",
    }

    pessoaExemplo := Pessoa{
        nome:     "João",
        idade:    30,
        endereco: enderecoExemplo,
    }

    imprimirEnderecoCompleto(pessoaExemplo)
}
