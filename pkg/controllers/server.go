package controllers

import (
	"fmt"
	"os"
	"os/exec"
)

// ServerController is used to manage NGINX virtual hosts
type ServerController struct{}

// NewServerController is used to create a new instance of the Servier Controller
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

// Create is used to create a new virtual host
// serverName will be used as the vitual host's servername
// hostName will be used for the proxy pass of the virtual host
func (c *ServerController) Create(serverName string, hostName string) {
	nginxRoot := getEnv("NGINX_ROOT", "/etc/nginx")
	fmt.Printf("creating %s -> %s...\n", serverName, hostName)
	sitesAvailablePath := fmt.Sprintf("%s/sites-available/%s", nginxRoot, serverName)
	sitesEnabledPath := fmt.Sprintf("%s/sites-enabled/%s", nginxRoot, serverName)
	f, err := os.Create(sitesAvailablePath)
	checkErr(err)
	f.Write([]byte(fmt.Sprintf(serverTemplate, serverName, hostName)))
	f.Sync()
	f.Close()
	_, err = exec.Command("ln", "-Fs", sitesAvailablePath, sitesEnabledPath).Output()
	checkErr(err)
}

const serverTemplate = "server {\n  listen 80;\n  listen [::]:80;\n  " +
	"\n  server_name %s;\n" +
	"\n  location / {\n" +
	"\n    proxy_pass %s;\n" +
	"\n  }\n\n}"
