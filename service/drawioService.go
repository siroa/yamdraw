package service

import (
	"generate/domain/drawio"
	"generate/domain/layout"
)

type DefaultDrawioService struct {
	Cs *drawio.MxCells
}

// MxCellsから最初のMxCell2つを生成する
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
		mcells, i = mcells.CreateGroup(p, v.PositionX, v.PositionY, v.Width, v.Height) // サーバーのグループ化
		mcells = mcells.CreateServer(v.ID, v.Name, i, v.Width, v.Height)
		mcells = setMws(i, mcells, v.Mws)
	}
	return mcells
}

func setMws(i string, mcells *DefaultDrawioService, lms []*layout.LayoutMw) *DefaultDrawioService {
	for _, lm := range lms {
		var ip string
		mcells, ip = mcells.CreateGroup(i, lm.PositionX, lm.PositionY, lm.Width, lm.Height) // MWのグループ化
		mcells = mcells.CreateMw(lm.ID, lm.Name, ip, lm.Width, lm.Height)
		for _, v := range lm.DB {
			mcells = mcells.CreateDB(v.ID, v.Name, ip, v.PositionX, v.PositionY, v.Width, v.Height)
		}
		for _, v := range lm.Process {
			mcells = mcells.CreateProcess(v.ID, v.Name, ip, v.PositionX, v.PositionY, v.Width, v.Height)
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

func (d *DefaultDrawioService) CreateGroup(parent string, x, y, w, h int) (*DefaultDrawioService, string) {
	var i string
	geo := drawio.NewGeometry().SetPosition(x, y).SetSize(w, h)
	d.Cs, i = d.Cs.NewGroup(parent, geo)
	return d, i
}

// 今は使わない
func (d *DefaultDrawioService) CreateCommunication(id, name string) *DefaultDrawioService {
	sourcePoint := drawio.NewPoint().SetPosition("160", "450").SetKind("sourcePoint")
	targetPoint := drawio.NewPoint().SetPosition("210", "400").SetKind("targetPoint")
	mps := []drawio.MxPoint{*sourcePoint, *targetPoint}
	geo := drawio.NewGeometryArrow(mps)
	d.Cs = d.Cs.NewCommunication(id, name, "4", "7", 0, 0, geo)
	return d
}

// 今は使わない
func (d *DefaultDrawioService) CreateReference(id, name string) *DefaultDrawioService {
	sourcePoint := drawio.NewPoint().SetPosition("70", "710").SetKind("sourcePoint")
	targetPoint := drawio.NewPoint().SetPosition("140", "600").SetKind("targetPoint")
	mps := []drawio.MxPoint{*sourcePoint, *targetPoint}
	geo := drawio.NewGeometryArrow(mps)
	d.Cs = d.Cs.NewReference(id, name, "7", "4", 0, 0, geo)
	return d
}

// simple通信
func (d *DefaultDrawioService) CreateSimpleCommnication(id, name, source, target string, entry, exit int) *DefaultDrawioService {
	geo := drawio.NewGeometrySimpleArrow()
	d.Cs = d.Cs.NewCommunication(id, name, source, target, entry, exit, geo)
	return d
}

// simple参照
func (d *DefaultDrawioService) CreateSimpleReference(id, name, source, target string, entry, exit int) *DefaultDrawioService {
	geo := drawio.NewGeometrySimpleArrow()
	d.Cs = d.Cs.NewReference(id, name, source, target, entry, exit, geo)
	return d
}
