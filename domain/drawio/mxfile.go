package drawio

import "encoding/xml"

type Mxfile struct {
	XMLName xml.Name `xml:"mxfile"`
	Host    string   `xml:"host,attr"`
	Diagram Diagram
}

func (f Mxfile) NewFile(dig Diagram) Mxfile {
	mf := Mxfile{
		Host:    "65bd71144e",
		Diagram: dig,
	}
	return mf
}
