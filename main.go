package main

import (
	"fmt"
	"github.com/LucasCarioca/nginx-reverse-proxy-cli/pkg/controllers"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version string
)

func main() {
	incorrectUsage := len(os.Args) < 2
	command := os.Args[1]
	if incorrectUsage || command == "help" {
		fmt.Println("USAGE: proxy [COMMAND]")
		fmt.Println("COMMANDS:")
		fmt.Println("\t create 	- used to create new virtual hosts")
		fmt.Println("\t version 	- check the current version of this cli")
		fmt.Println("\t help 		- see usage information")
		if incorrectUsage {
			os.Exit(0)
		}
	}
	if command == "create" {
		var serverName string
		var hostName string
		create := &cli.App{
			Name:  "create",
			Usage: "Create a new reverse proxy",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "name",
					Aliases:     []string{"n"},
					Value:       "example.com",
					Usage:       "Domain for the new server",
					Destination: &serverName,
				},
				&cli.StringFlag{
					Name:        "host",
					Aliases:     []string{"H"},
					Value:       "http://0.0.0.0",
					Usage:       "Host for the new server",
					Destination: &hostName,
				},
			},
			Action: func(c *cli.Context) error {
				servers := controllers.NewServerController()
				servers.Create(serverName, hostName)
				return nil
			},
		}
		err := create.Run(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}
	}
	if command == "version" {
		fmt.Printf("Current version: v%s\n", version)
	}
}
