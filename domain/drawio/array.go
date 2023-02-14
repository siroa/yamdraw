package drawio

import "encoding/xml"

type Array struct {
	XMLName xml.Name `xml:"Array"`
	As      string   `xml:"as,attr"`
	MxPoint []*MxPoint
}

func NewArray() *Array {
	mg := Array{
		As: "point",
	}
	return &mg
}

func (a *Array) SetMxPoint(mp *MxPoint) *Array {
	a.MxPoint = append(a.MxPoint, mp)
	return a
}
