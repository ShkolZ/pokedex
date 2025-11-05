package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("To Start using Pokedex start with 'help' command!")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		commandName := strings.ToLower(input[0])
		var arg string
		if len(input) > 1 {
			arg = strings.ToLower(input[1])
		}

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(cfg, arg)
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Printf("command '%v' does not exist\n", commandName)
		}

	}
}

func cleanInput(input string) []string {
	output := make([]string, 0)

	var str string
	for _, c := range input {
		if c != ' ' {
			str += string(c)
		} else if c == ' ' && len(str) != 0 {
			output = append(output, str)
			str = ""
		}

	}
	if str != "" {
		output = append(output, str)
	}

	return output
}
