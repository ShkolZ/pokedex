package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/ShkolZ/pokedexcli/internal/pokeapi"
	"github.com/ShkolZ/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	commandName string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			commandName: "help",
			description: "displays a help message",
			callback:    helpCommand,
		},
		"exit": {
			commandName: "exit",
			description: "exit from the pokedex",
			callback:    exitCommand,
		},
		"map": {
			commandName: "map",
			description: "shows next 20 locations of Pokemon world",
			callback:    mapCommand,
		},
		"mapb": {
			commandName: "mapf",
			description: "shows prev 20 locations of Pokemon world",
			callback:    mapbCommand,
		},
		"explore": {
			commandName: "explore",
			description: "shows all pokemons that can be encountered on certain area",
			callback:    exploreCommand,
		},
		"catch": {
			commandName: "catch",
			description: "let's you catch some pokemons",
			callback:    catchCommand,
		},
		"inspect": {
			commandName: "inspect",
			description: "let's you see details about pokemon you caught",
			callback:    inpectCommand,
		},
		"pokedex": {
			commandName: "pokedex",
			description: "shows all your pokemons",
			callback:    pokedexCommand,
		},
		"debug": {
			commandName: "debug",
			description: "shows values of config",
			callback:    debugCommand,
		},
	}
}

func exitCommand(cfg *config, arg string) error {
	fmt.Println("Exiting the Pokedex... Bye!")
	os.Exit(0)
	return nil
}

func helpCommand(cfg *config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n")

	for command, commandStruct := range getCommands() {
		fmt.Printf("%v - %v\n", command, commandStruct.description)
	}
	return nil
}

func mapCommand(cfg *config, arg string) error {
	locations := pokeapi.Locations{}

	ok, data := pokecache.CacheInst.Get(cfg.nextLocation)

	if !ok {

		var err error
		data, err := pokeapi.GetMaps(cfg.nextLocation)
		if err != nil {
			return err
		}

		pokecache.CacheInst.Add(cfg.nextLocation, data)

		if err := json.Unmarshal(data, &locations); err != nil {
			return err
		}

	} else {

		err := json.Unmarshal(data, &locations)
		if err != nil {
			return err
		}
	}
	if locations.Next != nil {
		cfg.nextLocation = *locations.Next
	}
	if locations.Previous != nil {
		cfg.prevLocation = *locations.Previous
	}

	fmt.Println("Found Areas: ")
	for _, loc := range locations.Results {
		fmt.Printf("- %v\n", loc.Name)
	}

	return nil
}

func mapbCommand(cfg *config, arg string) error {
	locations := pokeapi.Locations{}

	ok, data := pokecache.CacheInst.Get(cfg.prevLocation)

	if !ok {
		var err error
		data, err := pokeapi.GetMaps(cfg.prevLocation)
		if err != nil {
			fmt.Print(err)
		}

		pokecache.CacheInst.Add(cfg.prevLocation, data)

	} else {

		err := json.Unmarshal(data, &locations)
		if err != nil {
			return err
		}
	}
	if locations.Next != nil {
		cfg.nextLocation = *locations.Next
	}
	if locations.Previous != nil {
		cfg.prevLocation = *locations.Previous
	}
	fmt.Println("Found Areas: ")
	for _, loc := range locations.Results {
		fmt.Printf("- %v\n", loc.Name)
	}

	return nil
}

func exploreCommand(cfg *config, arg string) error {
	reqUrl := fmt.Sprintf("%v/%v", pokeapi.StartUrl, arg)
	pokemons := pokeapi.PokemonEncounters{}

	ok, data := pokecache.CacheInst.Get(reqUrl)
	if !ok {
		data, err := pokeapi.GetPokemons(reqUrl)
		if err != nil {
			return err
		}

		pokecache.CacheInst.Add(reqUrl, data)

		if err := json.Unmarshal(data, &pokemons); err != nil {
			return err
		}

	} else {
		if err := json.Unmarshal(data, &pokemons); err != nil {
			return err
		}
	}
	fmt.Println("Found Pokemons: ")
	for _, pokemonEncounter := range pokemons.PokemonEncounters {
		fmt.Printf("- %v\n", pokemonEncounter.Pokemon.Name)
	}
	return nil
}

func catchCommand(cfg *config, arg string) error {
	reqUrl := fmt.Sprintf("%v/pokemon/%v", pokeapi.BaseUrl, arg)
	pokemon := Pokemon{}

	ok, data := pokecache.CacheInst.Get(arg)

	if !ok {
		data, err := pokeapi.GetPokemon(reqUrl)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, &pokemon); err != nil {
			return err
		}

	} else {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return err
		}
	}
	fmt.Printf("Trying to catch %v!\n", pokemon.Name)

	basicChance := 0.55
	chance := basicChance / (float64(pokemon.BaseExperience) / 100)
	randNumber := rand.Float64()

	if randNumber <= chance || cfg.tryCount > 4 {
		AddPokemon(cfg, pokemon)
		cfg.tryCount = 0
	} else {
		fmt.Printf("Unlucky you weren't able to catch %v :(\n", pokemon.Name)
		if pokemon.Name != cfg.prevPokemon.Name {
			cfg.tryCount = 1
			cfg.prevPokemon = pokemon
			return nil
		}
		cfg.tryCount++
	}

	return nil
}

func pokedexCommand(cfg *config, arg string) error {
	GetPokemons(cfg)
	return nil
}

func inpectCommand(cfg *config, arg string) error {
	ShowInfo(cfg, arg)
	return nil
}

func debugCommand(cfg *config, arg string) error {
	fmt.Printf("%v\n", cfg.tryCount)
	fmt.Printf("%v\n", cfg.prevPokemon.Name)
	fmt.Printf("%v\n", cfg.nextLocation)
	fmt.Printf("%v\n", cfg.prevLocation)
	return nil
}

// catch clefairy
