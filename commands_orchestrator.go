package main

import (
	"github.com/mjh1207/gopokedex/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(conf *config, param string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next *string
	previous *string
	pokedex map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Get next 20 Pokemon area locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get previous 20 Pokemon area locations",
			callback: commandMapB,
		},
		"explore": {
			name: "explore",
			description: "Explore an area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch a Pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Show data for caught pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Display names of all caught pokemon",
			callback: commandPokedex,
		},
	}
}

