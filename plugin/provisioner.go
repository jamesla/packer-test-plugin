package james

import (
	"fmt"
	"github.com/hashicorp/packer/packer"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/template/interpolate"
)

type Provisioner struct {
	common.PackerConfig `mapstructure:",squash"`
	ctx                 interpolate.Context
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	fmt.Printf("in provisioner")
	return nil
}

func (p *Provisioner) Provision(ui packer.Ui, comm packer.Communicator) error {
	fmt.Printf("Provisioner")
	return nil
}

func (p *Provisioner) Cancel() {
	fmt.Printf("Cancelling")
}
