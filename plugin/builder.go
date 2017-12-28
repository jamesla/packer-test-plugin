package james

import (
	"fmt"
	openstack "github.com/hashicorp/packer/builder/openstack"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/packer/plugin"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/template/interpolate"
)

type Builder struct {
	common.PackerConfig `mapstructure:",squash"`
	ctx                 interpolate.Context
}

func (b *Builder) Cancel() {
	fmt.Printf("in cancel")
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	return nil, nil
}
func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	ui.Say("in say")
	server, _ := plugin.Server()
	server.RegisterBuilder(new(openstack.Builder))
	server.Serve()
	return nil, nil
}
