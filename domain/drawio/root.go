package drawio

import "encoding/xml"

type Root struct {
	XMLName xml.Name `xml:"root"`
	MxCells []MxCell
}

func (r Root) NewRoot(cells []MxCell) Root {
	ro := Root{
		MxCells: cells,
	}
	return ro
}

func (r *Root) AddRoot(cell MxCell) *Root {
	r.MxCells = append(r.MxCells, cell)
	return r
}
