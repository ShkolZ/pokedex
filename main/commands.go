package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ShkolZ/pokedexcli/internal/pokeapi"
	"github.com/ShkolZ/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	commandName string
	description string
	callback    func(*config) error
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
		"debug": {
			commandName: "debug",
			description: "shows values of config",
			callback:    debugCommand,
		},
	}
}

func exitCommand(cfg *config) error {
	fmt.Println("Exiting the Pokedex... Bye!")
	os.Exit(0)
	return nil
}

func helpCommand(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for command, commandStruct := range getCommands() {
		fmt.Printf("%v - %v\n", command, commandStruct.description)
	}
	return nil
}

func mapCommand(cfg *config) error {
	locations := pokeapi.Locations{}

	ok, data := pokecache.CacheInst.Get(cfg.nextLocation)

	if !ok {

		var err error
		err, locations = pokeapi.GetMaps(cfg.nextLocation)
		if err != nil {
			fmt.Print(err)
		}

		data, _ := json.Marshal(locations)
		pokecache.CacheInst.Add(cfg.nextLocation, data)

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

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func mapbCommand(cfg *config) error {
	locations := pokeapi.Locations{}

	ok, data := pokecache.CacheInst.Get(cfg.prevLocation)

	if !ok {
		var err error
		err, locations = pokeapi.GetMaps(cfg.prevLocation)
		if err != nil {
			fmt.Print(err)
		}

		data, _ := json.Marshal(locations)
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

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func debugCommand(cfg *config) error {
	return nil
}
