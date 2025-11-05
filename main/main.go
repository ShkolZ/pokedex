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
}

func main() {

	cfg := &config{
		nextLocation: pokeapi.StartUrl,
		pokedex:      make(map[string]Pokemon),
	}

	startRepl(cfg)
}
