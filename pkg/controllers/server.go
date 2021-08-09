package controllers

import (
	"fmt"
	"os"
	"os/exec"
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
	nginxRoot := getEnv("NGINX_ROOT", "/etc/nginx")
	fmt.Printf("creating %s -> %s...\n", serverName, hostName)
	sitesAvailablePath := fmt.Sprintf("%s/sites-available/%s", nginxRoot, serverName)
	sitesEnabledPath := fmt.Sprintf("%s/sites-enabled/%s", nginxRoot, serverName)
	f, err := os.Create(sitesAvailablePath)
	checkErr(err)
	f.Write([]byte(fmt.Sprintf(SERVER, serverName, hostName)))
	f.Sync()
	f.Close()
	_, err = exec.Command("ln", "-Fs", sitesAvailablePath, sitesEnabledPath).Output()
	checkErr(err)
}

const SERVER = "server {\n  listen 80;\n  listen [::]:80;\n  " +
	"\n  server_name %s;\n" +
	"\n  location / {\n" +
	"\n    proxy_pass %s;\n" +
	"\n  }\n\n}"
