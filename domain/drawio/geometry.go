package drawio

import (
	"encoding/xml"
	"strconv"
)

type MxGeometry struct {
	XMLName  xml.Name `xml:"mxGeometry"`
	X        string   `xml:"x,attr,omitempty"`
	Y        string   `xml:"y,attr,omitempty"`
	Width    string   `xml:"width,attr"`
	Height   string   `xml:"height,attr"`
	Relative string   `xml:"relative,attr,omitempty"`
	As       string   `xml:"as,attr"`
	Array    *Array
	// MxPoint  *[]MxPoint
}

func NewGeometry() *MxGeometry {
	mg := MxGeometry{
		X:      "",
		Y:      "",
		Width:  "",
		Height: "",
		As:     "geometry",
	}
	return &mg
}

func NewGeometrySimpleArrow() *MxGeometry {
	mg := MxGeometry{
		Relative: "1",
		As:       "geometry",
	}
	return &mg
}

// func NewGeometryArrow(mp []MxPoint) *MxGeometry {
// 	mg := MxGeometry{
// 		Width:    "50",
// 		Height:   "50",
// 		Relative: "1",
// 		As:       "geometry",
// 		MxPoint:  &mp,
// 	}
// 	return &mg
// }

func NewGeometryArrow() *MxGeometry {
	return &MxGeometry{
		Relative: "1",
		As:       "geometry",
		Array:    &Array{},
	}
}

func (g *MxGeometry) SetArray(array *Array) *MxGeometry {
	g.Array = array
	return g
}

func (g *MxGeometry) SetPosition(x, y int) *MxGeometry {
	g.X = toString(x)
	g.Y = toString(y)
	return g
}

func (g *MxGeometry) SetSize(width, height int) *MxGeometry {
	g.Width = toString(width)
	g.Height = toString(height)
	return g
}

func toString(v int) string {
	return strconv.Itoa(v)
}
