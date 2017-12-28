package main

import "fmt"

import (
	"github.com/hashicorp/packer/packer/plugin"
	serverspec "serverspec/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(serverspec.Builder))
	server.Serve()
}
