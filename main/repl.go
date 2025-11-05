package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func startRepl(cfg *config) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	// currentCommand := ""
	reader := bufio.NewReader(os.Stdin)
	currentCommand := ""
	temp := ""
	commandIndex := 0

	fmt.Println("To Start using Pokedex start with 'help' command!")
	ClearLine()

	for {
		commandIndex = len(cfg.commands)
		currentCommand = ""
		fmt.Print("Pokedex > ")
	innerLoop:
		for {
			term.MakeRaw(int(os.Stdin.Fd()))
			b, err := reader.ReadByte()
			if err != nil {
				break
			}
			switch b {
			case 3:
				ClearLine()
				term.Restore(int(os.Stdin.Fd()), oldState)
				os.Exit(0)
			case 27:
				_, err = reader.ReadByte()
				if err != nil {
					return
				}
				b, err = reader.ReadByte()
				if err != nil {
					return
				}
				if b == 65 {
					if commandIndex > 0 {
						commandIndex--
						ClearLine()
						fmt.Printf("Pokedex > %v", cfg.commands[commandIndex])
					}
				} else if b == 66 {
					if commandIndex < len(cfg.commands)-1 {
						commandIndex++
						ClearLine()
						fmt.Printf("Pokedex > %v", cfg.commands[commandIndex])
					}
				}
			case 127:
				if len(temp) > 0 {
					tempLen := len(temp) - 1
					temp = temp[:tempLen]
					ClearLine()
					fmt.Printf("Pokedex > %v", temp)
				}
			case 13:

				if commandIndex < len(cfg.commands) {
					currentCommand = cfg.commands[commandIndex]
					cfg.commands = append(cfg.commands, currentCommand)
				} else {
					currentCommand = temp
					cfg.commands = append(cfg.commands, currentCommand)
				}
				ClearLine()
				temp = ""
				term.Restore(int(os.Stdin.Fd()), oldState)
				fmt.Println()
				break innerLoop
			default:
				ClearLine()
				temp += string(b)
				fmt.Printf("Pokedex > %v", temp)
			}
		}
		input := cleanInput(currentCommand)
		commandName := strings.ToLower(input[0])
		arg := ""

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
