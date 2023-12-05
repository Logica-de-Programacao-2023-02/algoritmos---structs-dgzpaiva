package main

import (
	"fmt"
)

type Filme struct {
	titulo      string
	diretor     string
	ano         int
	avaliacoes  []int
}

func (f *Filme) adicionarAvaliacao(avaliacao int) {
	f.avaliacoes = append(f.avaliacoes, avaliacao)
}

func (f *Filme) removerAvaliacao(index int) {
	if index >= 0 && index < len(f.avaliacoes) {
		f.avaliacoes = append(f.avaliacoes[:index], f.avaliacoes[index+1:]...)
	}
}

func (f Filme) calcularMediaAvaliacoes() float64 {
	total := 0.0
	for _, avaliacao := range f.avaliacoes {
		total += float64(avaliacao)
	}
	if len(f.avaliacoes) > 0 {
		return total / float64(len(f.avaliacoes))
	}
	return 0.0
}

func (f Filme) imprimirInfo() {
	fmt.Printf("Título: %s\nDiretor: %s\nAno: %d\n", f.titulo, f.diretor, f.ano)
	media := f.calcularMediaAvaliacoes()
	fmt.Printf("Média das avaliações: %.2f\n", media)
}

func main() {
	filme := Filme{
		titulo:     "
