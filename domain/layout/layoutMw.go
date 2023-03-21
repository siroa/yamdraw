package layout

const (
	MWMARGIN      = 40
	MWHEIGHTSCALE = 1.5
	DEFAULTWIDHT  = 80
	DEFAULTHEIGHT = 100
)

type LayoutMw struct {
	ID          string
	Name        string
	Width       int
	Height      int
	PositionX   int
	PositionY   int
	Accessories []*LayoutAccessory
	KindNum     map[string]int
}

func CreateLayoutMw(id, name string) *LayoutMw {
	return &LayoutMw{
		ID:        id,
		Name:      name,
		Width:     0,
		Height:    0,
		PositionX: 0,
		PositionY: 0,
		KindNum:   make(map[string]int),
	}
}

// Calculate and set Mw width
func (d *LayoutMw) CalcWidth() {
	if len(d.Accessories) == 0 {
		d.Width = DEFAULTWIDHT
		return
	}
	longestWidth := 0
	currentKind := d.Accessories[0].Kind
	for _, v := range d.Accessories {
		if currentKind != v.Kind {
			currentKind = v.Kind
			if longestWidth > d.Width {
				d.Width = longestWidth
			}
			longestWidth = 0
		}
		longestWidth += v.Width + MWMARGIN
	}
	if longestWidth > d.Width {
		d.Width = longestWidth
	}
	if currentKind == d.Accessories[0].Kind {
		d.Width = longestWidth
	}
}

// Calculate and set the height of Mw
func (d *LayoutMw) CalcHeight() {
	if len(d.Accessories) == 0 {
		d.Height = DEFAULTHEIGHT
		return
	}
	currentKind := d.Accessories[0].Kind
	sumHeight := d.Accessories[0].Height + MARGIN
	for _, v := range d.Accessories {
		if currentKind != v.Kind {
			currentKind = v.Kind
			sumHeight += v.Height + MARGIN
			//d.Height = int(MWHEIGHTSCALE*float64(sumHeight)) + 20
			d.Height = 40 + sumHeight
		}
	}
	if currentKind == d.Accessories[0].Kind {
		//d.Height = int(MWHEIGHTSCALE*float64(d.Accessories[0].Height)) + 20
		d.Height = 40 + sumHeight
	}
}

func (d *LayoutMw) countAccessories() {
	if len(d.Accessories) == 0 {
		return
	}
	count := 0
	currentKind := d.Accessories[0].Kind
	for _, v := range d.Accessories {
		if currentKind != v.Kind {
			d.KindNum[currentKind] = count
			count = 0
			currentKind = v.Kind
		}
		count += 1
	}
	d.KindNum[currentKind] = count
}

func (d *LayoutMw) CalcAccessoriesPosion() {
	if len(d.Accessories) == 0 {
		return
	}
	d.countAccessories()
	currentKind := d.Accessories[0].Kind
	y := d.Height - (d.Accessories[0].Height + 10)

	padding := (d.Width - d.Accessories[0].Width*d.KindNum[currentKind]) / (d.KindNum[currentKind] + 1)
	x := padding
	for _, v := range d.Accessories {
		if currentKind != v.Kind {
			currentKind = v.Kind
			y -= v.Height + 20
			padding = (d.Width - v.Width*d.KindNum[currentKind]) / (d.KindNum[currentKind] + 1)
			x = padding
		}
		v.SetPosition(x, y)
		x += v.Width + padding
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
