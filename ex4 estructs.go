package main

import (
	"fmt"
	"time"
)

type Musica struct {
	titulo   string
	artista  string
	duracao  time.Duration
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func calcularDuracaoTotal(musicas []Musica) time.Duration {
	var duracaoTotal time.Duration
	for _, musica := range musicas {
		duracaoTotal += musica.duracao
	}
	return duracaoTotal
}

func imprimirPlaylistInfo(p Playlist) {
	fmt.Println("Nome da Playlist:", p.nome)
	duracaoTotal := calcularDuracaoTotal(p.musicas)
	fmt.Println("Duração Total da Playlist:", duracaoTotal)

	fmt.Println("Músicas na Playlist:")
	for _, musica := range p.musicas {
		fmt.Printf("Título: %s, Artista: %s, Duração: %s\n", musica.titulo, musica.artista, musica.duracao)
	}
}

func main() {
	musicas := []Musica{
		{titulo: "Música 1", artista: "Artista 1", duracao: 3 * time.Minute + 45 * time.Second},
		{titulo: "Música 2", artista: "Artista 2", duracao: 4 * time.Minute + 20 * time.Second},

	}

	minhaPlaylist := Playlist{nome: "Minha Playlist", musicas: musicas}
	imprimirPlaylistInfo(minhaPlaylist)
}
