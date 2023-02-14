package service

import (
	"generate/domain/drawio"
	"generate/domain/layout"
)

type DefaultDrawioService struct {
	Cs *drawio.MxCells
}

// Generate the first two MxCells from MxCells
func (d DefaultDrawioService) NewDrawioService() DefaultDrawioService {
	cells := drawio.MxCells{}
	service := DefaultDrawioService{}
	service.Cs = cells.InitCell()
	return service
}

func (d DefaultDrawioService) Build(service DefaultDrawioService, nls []*layout.LayoutNetwork, pls []layout.LayoutProcedure) drawio.Mxfile {
	var mcells *DefaultDrawioService
	for _, v := range nls {
		var p string
		mcells, p = service.CreateGroup("Generate1", v.PositionX, v.PositionY, v.Width, v.Height)
		mcells = mcells.CreateNetwork(v.ID, v.Name, p, v.Width, v.Height)
		mcells = setServer(p, mcells, v.Servers)
	}
	for _, v := range pls {
		if v.Kind == "comm" {
			mcells.CreateSimpleCommnication("", v.Route, v.Source, v.Target, v.Entry, v.Exit)
		} else if v.Kind == "ref" {
			mcells.CreateSimpleReference("", v.Route, v.Source, v.Target, v.Entry, v.Exit)
		} else if v.Kind == "update" {
			mcells.CreateSimpleCommnication("", v.Route, v.Source, v.Target, v.Entry, v.Exit)
		}
	}
	r := drawio.Root{}.NewRoot(mcells.Cs.Cells)
	gm := drawio.NewGraphModel(r)
	dig := drawio.Diagram{}.NewDiagram(gm)
	mf := drawio.Mxfile{}.NewFile(dig)
	return mf
}

func setServer(p string, mcells *DefaultDrawioService, ls []*layout.LayoutServer) *DefaultDrawioService {
	for _, v := range ls {
		var i string
		mcells, i = mcells.CreateGroup(p, v.PositionX, v.PositionY, v.Width, v.Height) // Server Grouping
		mcells = mcells.CreateServer(v.ID, v.Name, i, v.Width, v.Height)
		mcells = setMws(i, mcells, v.Mws)
	}
	return mcells
}

func setMws(i string, mcells *DefaultDrawioService, lms []*layout.LayoutMw) *DefaultDrawioService {
	for _, lm := range lms {
		var ip string
		mcells, ip = mcells.CreateGroup(i, lm.PositionX, lm.PositionY, lm.Width, lm.Height) // MW Grouping
		mcells = mcells.CreateMw(lm.ID, lm.Name, ip, lm.Width, lm.Height)
		for _, v := range lm.Accessories {
			switch k := v.Kind; {
			case k == "db":
				mcells = mcells.CreateDB(v.ID, v.Name, ip, v.PositionX, v.PositionY, v.Width, v.Height)
			case k == "process":
				mcells = mcells.CreateProcess(v.ID, v.Name, ip, v.PositionX, v.PositionY, v.Width, v.Height)
			case k == "doc":
				mcells = mcells.CreateDocument(v.ID, v.Name, ip, v.PositionX, v.PositionY, v.Width, v.Height)
			}

		}
	}
	return mcells
}

func (d *DefaultDrawioService) CreateNetwork(id, name, p string, w, h int) *DefaultDrawioService {
	geo := drawio.NewGeometry().SetSize(w, h)
	d.Cs = d.Cs.NewNetwork(id, name, p, geo)
	return d
}

func (d *DefaultDrawioService) CreateServer(id, name, p string, w, h int) *DefaultDrawioService {
	geo := drawio.NewGeometry().SetSize(w, h)
	d.Cs = d.Cs.NewServer(id, name, p, geo)
	return d
}

func (d *DefaultDrawioService) CreateMw(id, name, p string, w, h int) *DefaultDrawioService {
	geo := drawio.NewGeometry().SetSize(w, h)
	d.Cs = d.Cs.NewMw(id, name, p, geo)
	return d
}

func (d *DefaultDrawioService) CreateDB(id, name, p string, x, y, w, h int) *DefaultDrawioService {
	geo := drawio.NewGeometry().SetPosition(x, y).SetSize(w, h)
	d.Cs = d.Cs.NewDB(id, name, p, geo)
	return d
}

func (d *DefaultDrawioService) CreateProcess(id, name, p string, x, y, w, h int) *DefaultDrawioService {
	geo := drawio.NewGeometry().SetPosition(x, y).SetSize(w, h)
	d.Cs = d.Cs.NewProcess(id, name, p, geo)
	return d
}

func (d *DefaultDrawioService) CreateDocument(id, name, p string, x, y, w, h int) *DefaultDrawioService {
	geo := drawio.NewGeometry().SetPosition(x, y).SetSize(w, h)
	d.Cs = d.Cs.NewDocument(id, name, p, geo)
	return d
}

func (d *DefaultDrawioService) CreateGroup(parent string, x, y, w, h int) (*DefaultDrawioService, string) {
	var i string
	geo := drawio.NewGeometry().SetPosition(x, y).SetSize(w, h)
	d.Cs, i = d.Cs.NewGroup(parent, geo)
	return d, i
}

// I don't use it now.
func (d *DefaultDrawioService) CreateCommunication(id, name, source, target string, entry, exit int) *DefaultDrawioService {
	point := drawio.NewPoint().SetPosition("0", "0")
	array := drawio.NewArray().SetMxPoint(point)
	geo := drawio.NewGeometryArrow(array)
	d.Cs = d.Cs.NewCommunication(id, name, source, target, entry, exit, geo)
	return d
}

// I don't use it now.
// func (d *DefaultDrawioService) CreateReference(id, name string) *DefaultDrawioService {
// 	sourcePoint := drawio.NewPoint().SetPosition("70", "710").SetKind("sourcePoint")
// 	targetPoint := drawio.NewPoint().SetPosition("140", "600").SetKind("targetPoint")
// 	mps := []drawio.MxPoint{*sourcePoint, *targetPoint}
// 	geo := drawio.NewGeometryArrow(mps)
// 	d.Cs = d.Cs.NewReference(id, name, "7", "4", 0, 0, geo)
// 	return d
// }

// simple communication
func (d *DefaultDrawioService) CreateSimpleCommnication(id, name, source, target string, entry, exit int) *DefaultDrawioService {
	geo := drawio.NewGeometrySimpleArrow()
	d.Cs = d.Cs.NewCommunication(id, name, source, target, entry, exit, geo)
	return d
}

// simple communication
func (d *DefaultDrawioService) CreateSimpleReference(id, name, source, target string, entry, exit int) *DefaultDrawioService {
	geo := drawio.NewGeometrySimpleArrow()
	d.Cs = d.Cs.NewReference(id, name, source, target, entry, exit, geo)
	return d
}
