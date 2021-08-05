package controllers

import (
	"fmt"
	"os"
)

type ServerController struct{}

func NewServerController() ServerController {
	return ServerController{}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (c *ServerController) Create(serverName string, hostName string) {
	sitesAvailableRoot := os.Getenv("PROXY_SITES_AVAILABLE")
	fmt.Printf("creating %s -> %s...\n", serverName, hostName)
	f, err := os.Create(fmt.Sprintf("%s/%s", sitesAvailableRoot, serverName))
	check(err)
	defer f.Close()
	f.Write([]byte(fmt.Sprintf(SERVER, serverName, hostName)))
	f.Sync()
}

const SERVER = "server {\n  listen 80;\n  listen [::]:80;\n  " +
	"\n  server_name %s;\n" +
	"\n  location / {\n" +
	"\n    proxy_pass %s;\n" +
	"\n  }\n\n}"
