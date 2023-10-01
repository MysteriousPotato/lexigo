package generator

import (
	"github.com/dave/jennifer/jen"
)

type (
	field struct {
		Name string
		Type string
	}
	fields []field
)

func (f fields) toJen() []jen.Code {
	codes := make([]jen.Code, len(f))
	for i, f2 := range f {
		codes[i] = jen.Id(f2.Name).Id(f2.Type)
	}
	return codes
}
