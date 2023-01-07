package ui

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"internal"
)

// RunCLI runs the command-line interface for the "systems-environment-management" program
func RunCLI() {
	app := &cli.App{
		Name: "systems-environment-management",
		Usage: "Control and manage a system for growing bonsai trees in a controlled environment",
		Commands: []*cli.Command{
			{
				Name:    "start",
				Usage:   "Start the system",
				Action:  start,
			},
			{
				Name:    "stop",
				Usage:   "Stop the system",
				Action:  stop,
			},
			{
				Name:    "status",
				Usage:   "Get the status of the system",
				Action:  status,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	log.Println("Starting system...")
	// Start the system
	// ...
	return nil
}

func
