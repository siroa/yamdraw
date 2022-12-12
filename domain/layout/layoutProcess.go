package layout

const (
	PROWIDTH  int = 100
	PROHEIGHT int = 50
)

type LayoutProcess struct {
	ID        string
	Name      string
	Width     int
	Height    int
	PositionX int
	PositionY int
}

func CreateLayoutProcess(id, name string) *LayoutProcess {
	return &LayoutProcess{
		ID:        id,
		Name:      name,
		Width:     PROWIDTH,
		Height:    PROHEIGHT,
		PositionX: 0,
		PositionY: 0,
	}
}

func (d *LayoutProcess) SetSize(w, h int) {
	d.Width = w
	d.Height = h
}

func (d *LayoutProcess) SetPosition(x, y int) {
	d.PositionX = x
	d.PositionY = y
}
