package tfit

import (
	"fmt"
	"os"
)

type tfitter struct {
	config Config
}

type Tfitter interface {
	Tfit()
}

func (tf *tfitter) Tfit() {
	for _, module := range tf.config.Modules() {
		fmt.Printf("%s matched:\n%s\n", module.PrettyName(), module.Match(os.Stdin))
	}
}

func NewTfitter(c Config) Tfitter {
	return &tfitter{
		config: c,
	}
}
