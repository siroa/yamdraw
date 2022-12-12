package layout

const (
	SEVERWIDTH   = 100
	SERVERHEIGHT = 120
)

type LayoutServer struct {
	ID        string
	Name      string
	Width     int
	Height    int
	PositionX int
	PositionY int
	Mws       []*LayoutMw
}

func CreateLayoutServer(id, name string) *LayoutServer {
	return &LayoutServer{
		ID:        id,
		Name:      name,
		Width:     0,
		Height:    0,
		PositionX: 0,
		PositionY: 0,
	}
}

// ToDo: The more MW, the longer it will be horizontally, so we'll come up with an algorithm to determine the size later.
func (d *LayoutServer) CalcSize() {
	if len(d.Mws) == 0 {
		d.Width = SEVERWIDTH
		d.Height = SERVERHEIGHT
		return
	}
	sumWidth := 0
	maxHeight := 0
	for _, v := range d.Mws {
		sumWidth += v.Width + MWMARGIN
		if maxHeight <= v.Height {
			maxHeight = v.Height
		}
	}
	d.Width = sumWidth
	d.Height = int(float64(maxHeight) * MWHEIGHTSCALE)
}

func (d *LayoutServer) CalcMwPostion() {
	posX := MWMARGIN / 2
	for _, v := range d.Mws {
		posY := d.Height - (10 + v.Height)
		v.SetPosition(posX, posY)
		posX += (MWMARGIN / 2) + v.Width
	}
}

func (d *LayoutServer) SetSize(w, h int) {
	d.Width = w
	d.Height = h
}

func (d *LayoutServer) SetPosition(x, y int) {
	d.PositionX = x
	d.PositionY = y
}
