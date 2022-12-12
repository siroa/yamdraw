package layout

import (
	"errors"
)

type LayoutProcedure struct {
	Route  string
	Source string
	Target string
	Kind   string
	Entry  int
	Exit   int
}

const (
	COMMUNICATION = "->"
	REFERENCE     = ".>"
	UPDATE        = "-->"
)

func CreateLayoutProcedure(r, s, t, k string, en, ex int) (LayoutProcedure, error) {
	kind, err := selectKind(k)
	if err != nil {
		lp := LayoutProcedure{}
		return lp, err
	}
	return LayoutProcedure{
		Route:  r,
		Source: s,
		Target: t,
		Kind:   kind,
		Entry:  en,
		Exit:   ex,
	}, nil
}

func selectKind(k string) (string, error) {
	switch k {
	case COMMUNICATION:
		return "comm", nil
	case REFERENCE:
		return "ref", nil
	case UPDATE:
		return "update", nil
	default:
		return "", errors.New("Kind that does not exist.")
	}
}
