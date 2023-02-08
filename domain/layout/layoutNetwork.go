package layout

type LayoutNetwork struct {
	ID        string
	Name      string
	Width     int
	Height    int
	PositionX int
	PositionY int
	Servers   []*LayoutServer
	Procedure map[string]string
}

const (
	NETMARGIN = 20
	NETWIDTH  = 200
	NETHEIGHT = 250
)

func CreateLayoutNetwork(id, name string) *LayoutNetwork {
	return &LayoutNetwork{
		ID:        id,
		Name:      name,
		Width:     0,
		Height:    0,
		PositionX: 0,
		PositionY: 0,
	}
}

func (d *LayoutNetwork) CalcSize() {
	if len(d.Servers) == 0 {
		d.Width = NETWIDTH
		d.Height = NETHEIGHT
		return
	}
	sumWidth := 0
	comparedWidth := 0
	sumHeight := 0
	for _, v := range d.Servers {
		comparedWidth = v.Width + NETMARGIN
		if sumWidth < comparedWidth {
			sumWidth = comparedWidth
		}
		sumHeight += v.Height + NETMARGIN*2
	}
	d.Width = sumWidth
	d.Height = sumHeight
}

func (d *LayoutNetwork) CalcServerPosition() {
	posX := NETMARGIN / 2
	posY := d.Height
	for _, v := range d.Servers {
		posY -= (NETMARGIN / 2) + v.Height
		v.SetPosition(posX, posY)
	}
}

func (d *LayoutNetwork) SetSize(w, h int) {
	d.Width = w
	d.Height = h
}

func (d *LayoutNetwork) SetPosition(x, y int) {
	d.PositionX = x
	d.PositionY = y
}
