package main

import (
	"fmt"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
}

func AddPokemon(cfg *config, pokemon Pokemon) {
	fmt.Printf("You caught %v\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon
}

func GetPokemons(cfg *config) {
	fmt.Println("Your Pokemons: ")
	for key := range cfg.pokedex {
		fmt.Printf("- %v\n", key)
	}
}

func ShowInfo(cfg *config, pokemonName string) {
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Printf("You haven't caught %v yet :<\n", pokemonName)
		return
	}
	fmt.Printf("There you have information about %v!\n", pokemonName)
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weigth: %v\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("- %v\n", t.Type.Name)
	}
	fmt.Println("Abilities: ")
	for _, ab := range pokemon.Abilities {
		fmt.Printf("- %v\n", ab.Ability.Name)
	}
}
