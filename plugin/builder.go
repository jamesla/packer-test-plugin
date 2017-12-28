package james

import (
	"fmt"
	"github.com/hashicorp/packer/packer"
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
	fmt.Printf("in prepare")
	return nil, nil
}
func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	fmt.Printf("in run")
	return nil, nil
}
