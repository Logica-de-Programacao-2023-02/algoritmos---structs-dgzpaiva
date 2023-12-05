package main

import (
	"fmt"
	"time"
)

type Viagem struct {
	origem  string
	destino string
	data    time.Time
	preco   float64
}

func viagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		return Viagem{} 
	}

	viagemMaisCara := viagens[0]
	for _, viagem := range viagens {
		if viagem.preco > viagemMaisCara.preco {
			viagemMaisCara = viagem
		}
	}
	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{origem: "Cidade A", destino: "Cidade B", data: time.Now(), preco: 250.0},
		{origem: "Cidade C", destino: "Cidade D", data: time.Now(), preco: 180.0},
		{origem: "Cidade E", destino: "Cidade F", data: time.Now(), preco: 350.0},
	}

	viagemCara := viagemMaisCara(viagens)
	fmt.Println("A viagem mais cara é de", viagemCara.origem, "para", viagemCara.destino, "com preço de", viagemCara.preco)
}
