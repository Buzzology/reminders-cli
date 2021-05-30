package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Print("No command provided.")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":

		msg := "REMINDERS CLI - CLI BASICS"

		if len(os.Args) > 2 {

			// If provided it will use the --msg argument
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}

		fmt.Printf("Hello and welcome: %s\n", msg)

	case "help":
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
	}

	fmt.Println(os.Args)
}
