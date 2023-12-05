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

func playlistsComMusica(musicTitle string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == musicTitle {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}

func main() {
	musicas1 := []Musica{
		{titulo: "Música 1", artista: "Artista 1", duracao: 3 * time.Minute + 45 * time.Second},
		{titulo: "Música 2", artista: "Artista 2", duracao: 4 * time.Minute + 20 * time.Second},
	}

	musicas2 := []Musica{
		{titulo: "Música 3", artista: "Artista 3", duracao: 5 * time.Minute + 10 * time.Second},
		{titulo: "Música 1", artista: "Artista 1", duracao: 3 * time.Minute + 45 * time.Second},
	}

	playlist1 := Playlist{nome: "Playlist 1", musicas: musicas1}
	playlist2 := Playlist{nome: "Playlist 2", musicas: musicas2}

	tituloBuscado := "Música 1"
	playlists := []Playlist{playlist1, playlist2}

	playlistsEncontradas := playlistsComMusica(tituloBuscado, playlists)
	fmt.Printf("Playlists com a música '%s':\n", tituloBuscado)
	for _, playlist := range playlistsEncontradas {
		fmt.Println(playlist.nome)
	}
}
