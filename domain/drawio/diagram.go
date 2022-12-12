package drawio

import "encoding/xml"

type Diagram struct {
	XMLName      xml.Name `xml:"diagram"`
	Id           string   `xml:"id,attr"`
	Name         string   `xml:"name,attr"`
	MxGraphModel MxGraphModel
}

func (d Diagram) NewDiagram(mgm *MxGraphModel) Diagram {
	dig := Diagram{
		Id:           "DzRFVIfGvD17ysObrIGW",
		Name:         "Page-1",
		MxGraphModel: *mgm,
	}
	return dig
}
