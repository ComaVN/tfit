package tfitmodbase64

import (
	"encoding/base64"
	"io"
	"io/ioutil"
)

// Implements tfit.Module
type tfitmodbase64 struct {
}

func (m *tfitmodbase64) Match(r io.Reader) string {
	decoder := base64.NewDecoder(base64.StdEncoding, r)
	if b, err := ioutil.ReadAll(decoder); err == nil {
		return string(b)
	}
	return ""
}

func (m *tfitmodbase64) PrettyName() string {
	return "Base64"
}

func New() *tfitmodbase64 {
	return &tfitmodbase64{}
}
