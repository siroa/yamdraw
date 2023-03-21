package drawio

import "encoding/xml"

type MxGraphModel struct {
	XMLName    xml.Name `xml:"mxGraphModel"`
	Dx         string   `xml:"dx,attr"`
	Dy         string   `xml:"dy,attr"`
	Grid       string   `xml:"grid,attr"`
	GridSize   string   `xml:"gridSize,attr"`
	Guides     string   `xml:"guides,attr"`
	Tooltips   string   `xml:"tooltips,attr"`
	Connect    string   `xml:"connect,attr"`
	Arrows     string   `xml:"arrows,attr"`
	Fold       string   `xml:"fold,attr"`
	Page       string   `xml:"page,attr"`
	PageScale  string   `xml:"pageScale,attr"`
	PageWidth  string   `xml:"pageWidth,attr"`
	PageHeight string   `xml:"pageHeight,attr"`
	BackGround string   `xml:"background,attr"`
	Math       string   `xml:"math,attr"`
	Shadow     string   `xml:"shadow,attr"`
	Root       Root
}

func NewGraphModel(r Root) *MxGraphModel {
	mgm := MxGraphModel{
		Dx:         "439",
		Dy:         "828",
		Grid:       "1",
		GridSize:   "10",
		Guides:     "1",
		Tooltips:   "1",
		Connect:    "1",
		Arrows:     "1",
		Fold:       "1",
		Page:       "1",
		PageScale:  "1",
		PageWidth:  "827",
		PageHeight: "1169",
		BackGround: "#ffffff",
		Math:       "0",
		Shadow:     "0",
		Root:       r,
	}
	return &mgm
}

func (gm *MxGraphModel) SetPagePosition(dx, dy string) *MxGraphModel {
	gm.Dx = dx
	gm.Dy = dy
	return gm
}

func (gm *MxGraphModel) SetPageSize(width, height string) *MxGraphModel {
	gm.PageWidth = width
	gm.PageHeight = height
	return gm
}
