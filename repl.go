package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		input := scanner.Text()
		var param string
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		} else if len(words) > 1 {
			param = words[1]
		} else {
			param = ""
		}

		command, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(conf, param); err != nil {
			fmt.Printf("Unable to execute command: %v", err)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
