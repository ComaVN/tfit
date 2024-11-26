package tfit

import (
	"fmt"
	"io"
	"sync"

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
	reader io.ReadCloser
}

func (tf *tfitter) Tfit() {
	mr := multee.NewMulteeReader(tf.config.Input())
	moduleInstances := make([]moduleInstance, len(tf.config.Modules()))
	var wg sync.WaitGroup
	for idx, module := range tf.config.Modules() {
		moduleInstances[idx] = moduleInstance{module, mr.NewReader()}
	}
	for _, mi := range moduleInstances {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%s matched:\n%s\n", mi.module.PrettyName(), mi.module.Match(mi.reader))
			mi.reader.Close()
		}()
	}
	wg.Wait()
}

func NewTfitter(c Config) Tfitter {
	return &tfitter{
		config: c,
	}
}
