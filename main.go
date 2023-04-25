package main

import (
	"app/linuxsystem"
	"app/network"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Usage: ./app name password  http://serveraddress:port")
		return
	}

	name := strings.ReplaceAll(args[0], " ", "")
	password := strings.ReplaceAll(args[1], " ", "")
	network.ServAddr = strings.ReplaceAll(args[2], " ", "")

	token, err := network.GetAuthToken(name, password)
	if err != nil {
		fmt.Println("Error connecting to server", err)
		return
	}
	packages, err := linuxsystem.ListSnapPackages()
	if err != nil {
		fmt.Println("Error parsing snap package", err)
		return
	}

	err = network.AddSoftwares(name, token, packages)
	if err != nil {
		fmt.Println("Error sanding data to server", err)
		return
	}

	fmt.Println("Snap packages:")
	for _, pkg := range packages {
		fmt.Println(pkg)
	}
}
