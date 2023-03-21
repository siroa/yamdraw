package drawio

import "encoding/xml"

type MxPoint struct {
	XMLName xml.Name `xml:"mxPoint"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
}

func NewPoint() *MxPoint {
	mp := MxPoint{
		X: "",
		Y: "",
	}
	return &mp
}

func (p *MxPoint) SetPosition(x, y string) *MxPoint {
	p.X = x
	p.Y = y
	return p
}
