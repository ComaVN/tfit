package main

import (
	"github.com/ComaVN/tfit/internal/tfit"
)

func main() {
	config := tfit.NewConfig()
	tfitter := tfit.NewTfitter(config)
	tfitter.Tfit()
}
