package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/Buzzology/reminders-cli/client"
)


var (
	backendURIFlag = flag.String("backend", "http://localhost:8080", "Backend API URL")
	helpFlag       = flag.Bool("help", false, "Display a helper message.")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURIFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("Command switch error: %s", err)
		os.Exit(2)
	}
}