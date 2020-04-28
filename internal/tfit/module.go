package tfit

import (
	"io"
)

type module struct {
}

type Module interface {
	PrettyName() string
	Match(io.Reader) string
}
