package layout

type LayoutMw struct {
	ID        string
	Name      string
	Width     int
	Height    int
	PositionX int
	PositionY int
	DB        []*LayoutDB
	Process   []*LayoutProcess
}

const (
	MWMARGIN      = 40
	MWHEIGHTSCALE = 1.5
	DEFAULTWIDHT  = 80
	DEFAULTHEIGHT = 100
)

func CreateLayoutMw(id, name string) *LayoutMw {
	return &LayoutMw{
		ID:        id,
		Name:      name,
		Width:     0,
		Height:    0,
		PositionX: 0,
		PositionY: 0,
	}
}

// Mwのwidthを計算してセットする
func (d *LayoutMw) CalcWidth() {
	dbNum := len(d.DB)
	ProcessNum := len(d.Process)
	if len(d.DB) == 0 && len(d.Process) == 0 {
		d.Width = DBWIDTH + MARGIN
	} else if dbNum > ProcessNum {
		d.Width = dbNum * (DBHEIGHT + MARGIN)
	} else {
		d.Width = ProcessNum * (PROWIDTH + MARGIN)
	}
}

// Mwのheightを計算してセットする
func (d *LayoutMw) CalcHeight() {
	if len(d.DB) == 0 || len(d.Process) == 0 {
		d.Height = int(MWHEIGHTSCALE * float64(DBHEIGHT+20))
	} else {
		sumHeight := DBHEIGHT + PROHEIGHT + MARGIN
		d.Height = int(MWHEIGHTSCALE * float64(sumHeight))
	}
}

func (d *LayoutMw) CalcDBPostion() {
	dbnum := len(d.DB)
	processNum := len(d.Process)
	if dbnum == 0 {
		return
	}
	y := d.Height - (20 + DBHEIGHT + PROHEIGHT)
	if processNum == 0 {
		y = d.Height - (DBHEIGHT + 10)
	}
	padding := (d.Width - DBWIDTH*dbnum) / (dbnum + 1)
	posX := padding
	for _, v := range d.DB {
		v.SetPosition(posX, y)
		posX += DBWIDTH + padding
	}
}

func (d *LayoutMw) CalcProcessPostion() {
	processNum := len(d.Process)
	if processNum == 0 {
		return
	}
	y := d.Height - (10 + PROHEIGHT)
	padding := (d.Width - PROWIDTH*processNum) / (processNum + 1)
	posX := padding
	for _, v := range d.Process {
		v.SetPosition(posX, y)
		posX += PROWIDTH + padding
	}
}

func (d *LayoutMw) SetSize(w, h int) {
	d.Width = w
	d.Height = h
}

func (d *LayoutMw) SetPosition(x, y int) {
	d.PositionX = x
	d.PositionY = y
}
