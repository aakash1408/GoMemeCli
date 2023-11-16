package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Welcome > ")
		scanner.Scan()
		text := scanner.Text()

		words := cleanInput((text))
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		getCommands := availableCommands()

	}

}

type cliCommand struct {
	name         string
	description  string
	callfunction func()
}


func availableCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit" : {
		
		}
		"help" : {
		name : "help",
		description : "Helps the user by showing the avilable commands.",
		callfunction :,
	},
	
}
}
