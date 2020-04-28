package tfit

import (
	"io"
	"os"

	"github.com/ComaVN/tfit/internal/tfit/module/tfitmodbase64"
)

type config struct {
	input   io.Reader
	modules []Module
}

type Config interface {
	Input() io.Reader
	Modules() []Module
}

func (c *config) Input() io.Reader {
	return c.input
}

func (c *config) Modules() []Module {
	return c.modules
}

func NewConfig() *config {
	return &config{
		input: os.Stdin,
		modules: []Module{
			tfitmodbase64.New(),
		},
	}
}
