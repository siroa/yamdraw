package layout

const (
	DBWIDTH   int = 60
	DBHEIGHT  int = 50
	PROWIDTH  int = 100
	PROHEIGHT int = 50
	DOCWIDTH  int = 80
	DOCHEIGHT int = 50
	MARGIN    int = 20
)

type LayoutAccessory struct {
	ID        string
	Kind      string
	Name      string
	Width     int
	Height    int
	PositionX int
	PositionY int
}

func CreateLayoutAccessory(id, name, kind string) *LayoutAccessory {
	return &LayoutAccessory{
		ID:        id,
		Name:      name,
		Kind:      kind,
		Width:     0,
		Height:    0,
		PositionX: 0,
		PositionY: 0,
	}
}

func (d *LayoutAccessory) SetSize() *LayoutAccessory {
	switch k := d.Kind; {
	case k == "db":
		d.Width = DBWIDTH
		d.Height = DBHEIGHT
	case k == "process":
		d.Width = PROWIDTH
		d.Height = PROHEIGHT
	case k == "doc":
		d.Width = DOCWIDTH
		d.Height = DOCHEIGHT
	}
	return d
}

func (d *LayoutAccessory) SetPosition(x, y int) {
	d.PositionX = x
	d.PositionY = y
}
