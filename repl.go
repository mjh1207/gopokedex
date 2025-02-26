package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := &config {
		next: "",
		previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		
		command, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(conf); err != nil {
			fmt.Printf("Unable to execute command: %v", err)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
