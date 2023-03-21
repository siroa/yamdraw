package drawio

import "encoding/xml"

type Array struct {
	XMLName xml.Name `xml:"Array"`
	As      string   `xml:"as,attr"`
	MxPoint []*MxPoint
}

func NewArray() *Array {
	mg := Array{
		As: "points",
	}
	return &mg
}

func (a *Array) SetMxPoint(mp *MxPoint) {
	a.MxPoint = append(a.MxPoint, mp)
}
