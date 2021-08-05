package main

import (
	"github.com/LucasCarioca/nginx-reverse-proxy-cli/pkg/controllers"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var serverName string
	var hostName string
	app := &cli.App{
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
				Value:       "0.0.0.0",
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
