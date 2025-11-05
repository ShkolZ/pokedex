package main

import (
	"github.com/ShkolZ/pokedexcli/internal/pokeapi"
)

type config struct {
	nextLocation string
	prevLocation string
	pokedex      map[string]Pokemon
	prevPokemon  Pokemon
	tryCount     int
	commands     []string
}

type Node struct {
	Val  string
	Prev *Node
	Next *Node
}

func main() {

	cfg := &config{
		nextLocation: pokeapi.StartUrl,
		pokedex:      make(map[string]Pokemon),
		commands:     make([]string, 0),
	}

	startRepl(cfg)
}
