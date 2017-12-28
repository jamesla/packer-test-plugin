package main

import (
	"github.com/hashicorp/packer/packer/plugin"
	serverspec "serverspec/plugin"
)

func main() {
	server, _ := plugin.Server()
	server.RegisterBuilder(new(serverspec.Builder))
	server.Serve()
}
