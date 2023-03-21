package service

import (
	"errors"
	"fmt"
	"generate/domain/drawio"
	"generate/domain/layout"
	"math"
	"os"
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

// ToDo: need refactoring
func AllocateContents(nls []*layout.LayoutNetwork) []layout.LayoutContent {
	var lcs []layout.LayoutContent
	for _, n := range nls {
		lc := layout.NewLayoutContent(n.ID, 1, n.Width, n.Height, n.PositionX, n.PositionY)
		lcs = append(lcs, lc)
		for _, s := range n.Servers {
			sx := n.PositionX + s.PositionX
			sy := n.PositionY + s.PositionY
			if sy == 45 {
				fmt.Printf("sy: %d, a.PositionY: %d\n", sy, s.PositionY)
			}
			lc := layout.NewLayoutContent(s.ID, 2, s.Width, s.Height, sx, sy)
			lcs = append(lcs, lc)
			for _, m := range s.Mws {
				mx := sx + m.PositionX
				my := sy + m.PositionY
				lc := layout.NewLayoutContent(m.ID, 3, m.Width, m.Height, mx, my)
				lcs = append(lcs, lc)
				for _, a := range m.Accessories {
					ax := mx + a.PositionX
					ay := my + a.PositionY
					if ay == 54 {
						fmt.Printf("my: %d, a.PositionY: %d\n", my, a.PositionY)
					}
					lc := layout.NewLayoutContent(a.ID, 4, a.Width, a.Height, ax, ay)
					lcs = append(lcs, lc)
				}
			}
		}
	}
	return lcs
}

// temp func
func writeByres(arr [][]int) {
	// 出力先のファイルを作成する
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 二重配列の要素をテキストファイルに書き込む
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Fprintf(file, "%d ", arr[j][i])
		}
		fmt.Fprintln(file, "") // 改行を挿入する
	}
}

// ToDo: need refactoring
func AllocateField(lcs []layout.LayoutContent) layout.BackGround {
	b := layout.BackGround{}
	b.InitGround()
	for _, v := range lcs {
		w := v.Width / 10
		h := v.Height / 10
		x := v.PositionX / 10
		y := v.PositionY / 10
		fmt.Printf("type: %d, width: %d, height: %d, x: %d, y: %d\n", v.Type, w, h, x, y)
		for i := x; i < (x + w); i++ {
			for j := y; j < (y + h); j++ {
				if b.Field[i][j] < v.Type {
					b.Field[i][j] = v.Type
				}
			}
		}
	}
	writeByres(b.Field)
	return b
}

// build
func (d DefaultDrawioService) Build(service DefaultDrawioService, nls []*layout.LayoutNetwork, pls []layout.LayoutProcedure) drawio.Mxfile {
	lcs := AllocateContents(nls)
	b := AllocateField(lcs)
	var mcells *DefaultDrawioService
	for _, v := range nls {
		var p string
		mcells, p = service.CreateGroup("Generate1", v.PositionX, v.PositionY, v.Width, v.Height)
		mcells = mcells.CreateNetwork(v.ID, v.Name, p, v.Width, v.Height)
		mcells = setServer(p, mcells, v.Servers)
	}
	mcells = addRelationLine(mcells, pls, b, lcs)
	r := drawio.Root{}.NewRoot(mcells.Cs.Cells)
	gm := drawio.NewGraphModel(r)
	dig := drawio.Diagram{}.NewDiagram(gm)
	mf := drawio.Mxfile{}.NewFile(dig)
	return mf
}

// ToDo: need refactoring
func selectRoute(source, target string, bg layout.BackGround, lcs []layout.LayoutContent) (int, int) {
	var sContent layout.LayoutContent
	var tContent layout.LayoutContent
	// つなげるコンテントの抽出
	for _, v := range lcs {
		if v.ID == source {
			sContent = v
		}
		if v.ID == target {
			tContent = v
		}
	}
	if sContent.Type == 0 || tContent.Type == 0 {
		// Todo: need error handling
		fmt.Println("Not found content.")
	}
	sX := sContent.PositionX / 10
	sY := sContent.PositionY / 10
	tX := tContent.PositionX / 10
	tY := tContent.PositionY / 10

	// ソースとターゲット間のルートにおける最小値を確認する
	min := checkBaseValue(bg, sX, sY, tX, tY)
	fmt.Printf("min: %d\n", min)

	sWidth := (sContent.Width / 10)
	sHight := (sContent.Height / 10)
	xDistance := math.Abs(float64(tX - sX))
	yDistance := math.Abs(float64(tY - sY))
	fmt.Printf("sWidth: %d, sHight: %d, xd: %f, yd: %f\n", sWidth, sHight, xDistance, yDistance)
	wNear := isNear(sWidth, int(xDistance))
	hNear := isNear(sHight, int(yDistance))

	gapX := sX - tX
	gapY := sY - tY
	rs := calcPositionalRelationships(gapX)
	rs += calcPositionalRelationships(gapY)
	fmt.Printf("rs: %s\n", rs)
	exit, entry, _ := routing(rs, bg, wNear, hNear)
	fmt.Printf("exit: %d, entry: %d\n", exit, entry)
	return exit, entry
}

func escapeCoordinate(x, y, base int, bg layout.BackGround) (int, int) {
	value := 1000
	i := 1
	f := bg.Field[x][y]
	var lx, ly int
	for base <= value {
		value = bg.Field[x-i][y]
		lx = x - i
		i++
		if f < value {
			lx = 0
			ly = 0
			break
		}
	}
	return lx, ly
}

func isNear(s, d int) bool {
	if d > s {
		return false
	} else {
		return true
	}
}

type branch struct {
	x int
	y int
}

func checkBaseValue(bg layout.BackGround, sX, sY, tX, tY int) int {
	var minX int
	var minY int
	x := tX - sX
	y := tY - sY
	if x > 0 {
		minX = returnXmin(sX, tX, tY, bg)
	} else {
		minX = returnXmin(tX, sX, tY, bg)
	}
	if y > 0 {
		minY = returnYmin(sY, tX, tY, bg)
	} else {
		minY = returnYmin(tY, tX, sY, bg)
	}
	if minX < minY {
		return minX
	}
	return minY
}

func returnXmin(idx, x, y int, bg layout.BackGround) int {
	min := 100
	for i := idx; i <= x; i++ {
		if min > bg.Field[i][y] {
			min = bg.Field[i][y]
		}
	}
	return min
}

func returnYmin(idx, x, y int, bg layout.BackGround) int {
	min := 100
	for i := idx; i <= y; i++ {
		if min > bg.Field[x][i] {
			min = bg.Field[x][i]
		}
	}
	return min
}

// sX: 中間, sY: 中間, tX: 中間, tY: 中間
func searchBranch(bg layout.BackGround, sX, sY, tX, tY int) {
	current := 0
	pre := 100
	sameCount := 0

	var branches []branch
	x := sX
	y := sY
	direction := "right"
	for x != tX && y != tY {
		if direction == "right" {
			x += 1
		} else {
			y += 1
		}

		current = bg.Field[x][y]
		if current > pre {
			jump := sameCount / 2
			var b branch
			if direction == "right" {
				b = branch{x: x - jump, y: y}
				direction = "bottom"
			} else {
				b = branch{x: x, y: y - jump}
				direction = "right"
			}
			branches = append(branches, b)
			sameCount = 0
		} else if pre == current {
			sameCount += 1
		}
		pre = current
	}
}

// return exit, entry
func routing(rs string, bg layout.BackGround, wn, hn bool) (int, int, error) {
	switch rs {
	case "11": // 右下
		// ToDo
		return 1, 0, nil
	case "12": // 右
		// ToDo
		return 1, 3, nil
	case "13": // 右上
		// ToDo
		return 1, 2, nil
	case "21": // 下
		// ToDo
		return 2, 0, nil
	case "23": // 上
		// ToDo
		return 0, 2, nil
	case "31": // 左下
		// ToDo
		return 3, 1, nil
	case "32": // 左
		// ToDo
		return 3, 1, nil
	case "33": // 左上
		// ToDo
		return 3, 2, nil
	default:
		return 0, 0, errors.New("hogehoge")
	}
}

func calcPositionalRelationships(gap int) string {
	if gap < 0 {
		// sourceが左側でtargetが右側にいる
		// sourceが上側でtargetが下側にいる
		return "1"
	} else if gap == 0 {
		// sourceとtargetが同じ幅にいる
		// sourceとtargetが同じ高さにいる
		return "2"
	} else {
		// sourceが右側でtargetが左側にいる
		// sourceが下側でtargetが上側にいる
		return "3"
	}
}

func addRelationLine(mcells *DefaultDrawioService, pls []layout.LayoutProcedure, bg layout.BackGround, lcs []layout.LayoutContent) *DefaultDrawioService {
	for _, v := range pls {
		if v.Kind == "comm" {
			exit, entry := selectRoute(v.Source, v.Target, bg, lcs)
			mcells.CreateCommunication("", v.Route, v.Source, v.Target, entry, exit)
		} else if v.Kind == "ref" {
			mcells.CreateSimpleReference("", v.Route, v.Source, v.Target, v.Entry, v.Exit)
		} else if v.Kind == "update" {
			mcells.CreateSimpleCommnication("", v.Route, v.Source, v.Target, v.Entry, v.Exit)
		}
	}
	return mcells
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
	// source, target, entry, exit,
	array := drawio.NewArray()
	// for //
	point := drawio.NewPoint().SetPosition("440", "390")
	array.SetMxPoint(point)
	point = drawio.NewPoint().SetPosition("440", "630")
	array.SetMxPoint(point)
	// for //
	geo := drawio.NewGeometryArrow().SetArray(array)
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
