package main

import (
	"github.com/ShkolZ/pokedexcli/internal/pokeapi"
)

type config struct {
	nextLocation string
	prevLocation string
}

func main() {

	cfg := &config{
		nextLocation: pokeapi.StartUrl,
	}

	startRepl(cfg)
}
