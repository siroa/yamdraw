package layout

const (
	DBWIDTH   int = 60
	DBHEIGHT  int = 50
	PROWIDTH  int = 100
	PROHEIGHT int = 50
	MARGIN    int = 20
)

type LayoutAccessory struct {
	ID        string
	Name      string
	Width     int
	Height    int
	PositionX int
	PositionY int
}

func CreateLayoutAccessory(id, name string, w, h int) *LayoutAccessory {
	return &LayoutAccessory{
		ID:        id,
		Name:      name,
		Width:     w,
		Height:    h,
		PositionX: 0,
		PositionY: 0,
	}
}

func (d *LayoutAccessory) SetSize(w, h int) {
	d.Width = w
	d.Height = h
}

func (d *LayoutAccessory) SetPosition(x, y int) {
	d.PositionX = x
	d.PositionY = y
}
