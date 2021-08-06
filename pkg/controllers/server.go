package controllers

import (
	"fmt"
	"os"
)

type ServerController struct{}

func NewServerController() ServerController {
	return ServerController{}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func (c *ServerController) Create(serverName string, hostName string) {
	sitesAvailableRoot := getEnv("PROXY_SITES_AVAILABLE", "/etc/nginx/sites-available")
	fmt.Printf("creating %s -> %s...\n", serverName, hostName)
	f, err := os.Create(fmt.Sprintf("%s/%s", sitesAvailableRoot, serverName))
	checkErr(err)
	defer f.Close()
	f.Write([]byte(fmt.Sprintf(SERVER, serverName, hostName)))
	f.Sync()
}

const SERVER = "server {\n  listen 80;\n  listen [::]:80;\n  " +
	"\n  server_name %s;\n" +
	"\n  location / {\n" +
	"\n    proxy_pass %s;\n" +
	"\n  }\n\n}"
