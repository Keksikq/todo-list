package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo-list-for-levus/cmd"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Todo CLI!")
	fmt.Println("Type 'help' for available commands or 'exit' to quit.")

	for {
		fmt.Print("todo> ")
		if !reader.Scan() {
			break
		}
		input := strings.TrimSpace(reader.Text())

		if input == "" {
			continue
		}

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		if parts[0] == "help" {
			cmd.RootCmd.SetArgs([]string{"--help"})
			if err := cmd.RootCmd.Execute(); err != nil {
				fmt.Println("Error:", err)
			}
			continue
		}

		cmd.RootCmd.SetArgs(parts)

		if err := cmd.RootCmd.Execute(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
