package tfit

import (
	"fmt"
	"io"
	"os"

	"github.com/ComaVN/multee"
)

type tfitter struct {
	config Config
}

type Tfitter interface {
	Tfit()
}

type moduleInstance struct {
	module Module
	reader io.Reader
}

func (tf *tfitter) Tfit() {
	mr := multee.NewMulteeReader(tf.config.Input())
	moduleInstances := make([]moduleInstance, len(tf.config.Modules()))
	for idx, module := range tf.config.Modules() {
		moduleInstances[idx] = moduleInstance{module, mr.NewReader()}
	}
	for _, mi := range moduleInstances {
		fmt.Printf("%s matched:\n%s\n", mi.module.PrettyName(), mi.module.Match(os.Stdin))
	}
}

func NewTfitter(c Config) Tfitter {
	return &tfitter{
		config: c,
	}
}
