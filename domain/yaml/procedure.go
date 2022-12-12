package yaml

import (
	"reflect"
)

type Procedure struct {
	A string `yaml:"a,omitempty"`
	B string `yaml:"b,omitempty"`
	C string `yaml:"c,omitempty"`
	D string `yaml:"d,omitempty"`
	E string `yaml:"e,omitempty"`
	F string `yaml:"f,omitempty"`
	G string `yaml:"g,omitempty"`
	H string `yaml:"h,omitempty"`
	I string `yaml:"i,omitempty"`
	J string `yaml:"j,omitempty"`
	K string `yaml:"k,omitempty"`
	L string `yaml:"l,omitempty"`
	M string `yaml:"m,omitempty"`
	N string `yaml:"n,omitempty"`
	O string `yaml:"o,omitempty"`
	P string `yaml:"p,omitempty"`
	Q string `yaml:"q,omitempty"`
	R string `yaml:"r,omitempty"`
	S string `yaml:"s,omitempty"`
	T string `yaml:"t,omitempty"`
	U string `yaml:"u,omitempty"`
	V string `yaml:"v,omitempty"`
	W string `yaml:"w,omitempty"`
	X string `yaml:"x,omitempty"`
	Y string `yaml:"y,omitempty"`
	Z string `yaml:"z,omitempty"`
}

func (p Procedure) ToProcedureList() map[string]string {
	procedure := map[string]string{}

	value := reflect.ValueOf(p)
	types := reflect.TypeOf(p)

	numField := types.NumField()
	for i := 0; i < numField; i++ {
		field := types.Field(i)
		fieldValue := value.Field(i).String()
		if fieldValue != "" {
			procedure[field.Name] = fieldValue
		}
	}
	return procedure
}
