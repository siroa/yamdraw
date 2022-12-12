package layout

const (
	DBWIDTH  int = 60
	DBHEIGHT int = 50
	MARGIN   int = 20
)

type LayoutDB struct {
	ID        string
	Name      string
	Width     int
	Height    int
	PositionX int
	PositionY int
}

func CreateLayoutDB(id, name string) *LayoutDB {
	return &LayoutDB{
		ID:        id,
		Name:      name,
		Width:     DBWIDTH,
		Height:    DBHEIGHT,
		PositionX: 0,
		PositionY: 0,
	}
}

func (d *LayoutDB) SetSize(w, h int) {
	d.Width = w
	d.Height = h
}

func (d *LayoutDB) SetPosition(x, y int) {
	d.PositionX = x
	d.PositionY = y
}
