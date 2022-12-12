package drawio

import (
	"encoding/xml"
	"errors"
	"generate/utils/logger"
	"strconv"
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

type MxCell struct {
	XMLName     xml.Name `xml:"mxCell"`
	Id          string   `xml:"id,attr"`
	Value       string   `xml:"value,attr,omitempty"`
	Style       string   `xml:"style,attr,omitempty"`
	Connectable string   `xml:"connectable,attr,omitempty"`
	Vertex      string   `xml:"vertex,attr,omitempty"`
	Parent      string   `xml:"parent,attr,omitempty"`
	Source      string   `xml:"source,attr,omitempty"`
	Target      string   `xml:"target,attr,omitempty"`
	Edge        string   `xml:"edge,attr,omitempty"`
	MxGeometry  *MxGeometry
}

type MxCells struct {
	Cells []MxCell
	Count int
}

func (cs *MxCells) InitCell() *MxCells {
	mc0 := MxCell{
		Id: "Generate0",
	}
	mc1 := MxCell{
		Id:     "Generate1",
		Parent: "Generate0",
	}
	cs.Count = 1
	cs.Cells = append(cs.Cells, mc0)
	cs.Cells = append(cs.Cells, mc1)
	return cs
}

func (cs *MxCells) NewNetwork(id, name, parent string, mg *MxGeometry) *MxCells {
	cs.Count++
	i := setId(id, cs.Count)
	mc := MxCell{
		Id:         i,
		Value:      name,
		Style:      "whiteSpace=wrap;html=1;aspect=fixed;strokeColor=#000000;strokeWidth=1;fillColor=#B3B3B3;gradientColor=none;align=left;fontColor=#000000;verticalAlign=top;",
		Vertex:     "1",
		Parent:     parent,
		MxGeometry: mg,
	}
	cs.Cells = append(cs.Cells, mc)
	return cs
}

func (cs *MxCells) NewServer(id, name, parent string, mg *MxGeometry) *MxCells {
	cs.Count++
	i := setId(id, cs.Count)
	mc := MxCell{
		Id:         i,
		Value:      name,
		Style:      "whiteSpace=wrap;html=1;aspect=fixed;strokeColor=#000000;strokeWidth=1;fillColor=#f5f5f5;gradientColor=none;align=left;fontColor=#000000;verticalAlign=top;",
		Vertex:     "1",
		Parent:     parent,
		MxGeometry: mg,
	}
	cs.Cells = append(cs.Cells, mc)
	return cs
}

func (cs *MxCells) NewDB(id, name, parent string, mg *MxGeometry) *MxCells {
	cs.Count++
	i := setId(id, cs.Count)
	mc := MxCell{
		Id:         i,
		Value:      name,
		Style:      "shape=cylinder3;whiteSpace=wrap;html=1;boundedLbl=1;backgroundOutline=1;size=8;fontColor=#333333;fillColor=#f5f5f5;strokeColor=#666666;",
		Vertex:     "1",
		Parent:     parent,
		MxGeometry: mg,
	}
	cs.Cells = append(cs.Cells, mc)
	return cs
}

func (cs *MxCells) NewMw(id, name, parent string, mg *MxGeometry) *MxCells {
	cs.Count++
	i := setId(id, cs.Count)
	mc := MxCell{
		Id:         i,
		Value:      name,
		Style:      "rounded=1;whiteSpace=wrap;html=1;fillColor=#66B2FF;fontColor=#000000;strokeColor=#001DBC;align=center;verticalAlign=top;",
		Vertex:     "1",
		Parent:     parent,
		MxGeometry: mg,
	}
	cs.Cells = append(cs.Cells, mc)
	return cs
}

func (cs *MxCells) NewProcess(id, name, parent string, mg *MxGeometry) *MxCells {
	cs.Count++
	i := setId(id, cs.Count)
	mc := MxCell{
		Id:         i,
		Value:      name,
		Style:      "shape=process;whiteSpace=wrap;html=1;backgroundOutline=1;fontColor=#333333;fillColor=#f5f5f5;strokeColor=#666666;",
		Vertex:     "1",
		Parent:     parent,
		MxGeometry: mg,
	}
	cs.Cells = append(cs.Cells, mc)
	return cs
}

func (cs *MxCells) NewCommunication(id, name, source, target string, entry, exit int, mg *MxGeometry) *MxCells {
	cs.Count++
	i := setId(id, cs.Count)
	en, err := allocateEntryPos(entry)
	if err != nil {
		logger.Error(err.Error())
	}
	ex, err := allocateExitPos(exit)
	if err != nil {
		logger.Error(err.Error())
	}
	mc := MxCell{
		Id:         i,
		Value:      name,
		Style:      "endArrow=classic;html=1;strokeColor=#000000;entryDx=0;entryDy=0;entryPerimeter=0;labelBackgroundColor=#FFFF66;fontColor=#000000;" + en + ex,
		Edge:       "1",
		Parent:     "Generate1",
		Source:     source,
		Target:     target,
		MxGeometry: mg,
	}
	cs.Cells = append(cs.Cells, mc)
	return cs
}

func (cs *MxCells) NewReference(id, name, source, target string, entry, exit int, mg *MxGeometry) *MxCells {
	cs.Count++
	i := setId(id, cs.Count)
	en, err := allocateEntryPos(entry)
	if err != nil {
		logger.Error(err.Error())
	}
	ex, err := allocateExitPos(exit)
	if err != nil {
		logger.Error(err.Error())
	}
	mc := MxCell{
		Id:         i,
		Value:      name,
		Style:      "endArrow=classic;dashed=1;html=1;dashPattern=1 3;strokeWidth=2;fontColor=#000000;strokeColor=#000000;startArrow=none;startFill=0;entryDx=0;entryDy=0;entryPerimeter=0;exitX=0.75;exitY=1;exitDx=0;exitDy=0;endFill=1;labelBackgroundColor=#FFFF66;" + en + ex,
		Edge:       "1",
		Parent:     "Generate1",
		Source:     source,
		Target:     target,
		MxGeometry: mg,
	}
	cs.Cells = append(cs.Cells, mc)
	return cs
}

func (cs *MxCells) NewGroup(parent string, mg *MxGeometry) (*MxCells, string) {
	cs.Count++
	i := "Generate" + strconv.Itoa(cs.Count)
	mc := MxCell{
		Id:          i,
		Value:       "",
		Style:       "group",
		Vertex:      "1",
		Connectable: "0",
		Parent:      parent,
		MxGeometry:  mg,
	}
	cs.Cells = append(cs.Cells, mc)
	return cs, i
}

func allocateEntryPos(pos int) (string, error) {
	switch pos {
	case 0:
		return "entryX=0.5;entryY=0;", nil
	case 1:
		return "entryX=1;entryY=0.5;", nil
	case 2:
		return "entryX=0.5;entryY=1;", nil
	case 3:
		return "entryX=0;entryY=0.5;", nil
	default:
		return "entryX=0.5;entryY=0;", errors.New("not exist postion num. decide default entry.")
	}
}

func allocateExitPos(pos int) (string, error) {
	switch pos {
	case 0:
		return "exitX=0.5;exitY=0;", nil
	case 1:
		return "exitX=1;exitY=0.5;", nil
	case 2:
		return "exitX=0.5;exitY=1;", nil
	case 3:
		return "exitX=0;exitY=0.5;", nil
	default:
		return "exitX=0.5;exitY=0;", errors.New("not exist postion num. decide default entry.")
	}
}

func setId(id string, count int) string {
	if id == "" || id == "0" {
		return "Generate" + strconv.Itoa(count)
	} else {
		return id
	}
}
