package main

type cliCommand struct {
	name string
	description string
	callback func() error
}

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
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
			description: "",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "",
			callback: commandMapB,
		},
	}
}

