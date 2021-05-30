package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Print("No command provided.")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":

		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI BASICS - REMINDERS CLI", "The message for the greet command.")
		err := greetCmd.Parse(os.Args[2:]) // Index of 2 and up
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("Hello and welcome: %s\n", *msgFlag)

	case "help":
		fmt.Println("Some help message.")

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
	}

	fmt.Println(os.Args)
}
