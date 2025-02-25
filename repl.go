package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		
		command, ok := supportedCommands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Printf("Unable to execute command: %v", err)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
