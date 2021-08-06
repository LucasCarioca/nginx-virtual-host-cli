package main

import (
	"github.com/LucasCarioca/nginx-reverse-proxy-cli/pkg/controllers"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version   string
)

func main() {
	var serverName string
	var checkVersionFlag bool
	var hostName string
	app := &cli.App{
		Name:  "create",
		Usage: "Create a new reverse proxy",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "version",
				Usage:       "Get current cli version",
				Destination: &checkVersionFlag,
			},
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
			if checkVersionFlag {
				println(version)
				os.Exit(0)
			}
			servers := controllers.NewServerController()
			servers.Create(serverName, hostName)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
