package layout

const (
	WIDTH        = 500
	HEIGHT       = 500
	LAYOUTMERGIN = 40
)

type Layout struct {
	Actor          []*LayoutNetwork
	Main           []*LayoutNetwork
	Others         []*LayoutNetwork
	ActorWidth     int
	MainWidth      int
	OthersWidth    int
	LayoutContents []LayoutContent
}

type LayoutContent struct {
	ID        string
	Type      int
	Width     int
	Height    int
	PositionX int
	PositionY int
}

type BackGround struct {
	Field [][]int
}

func NewLayoutContent(id string, t, w, h, x, y int) LayoutContent {
	return LayoutContent{
		ID:        id,
		Type:      t,
		Width:     w,
		Height:    h,
		PositionX: x,
		PositionY: y,
	}
}

// 1 = 10px
// initial value is 0
func (b *BackGround) InitGround() {
	b.Field = make([][]int, WIDTH)
	for i := 0; i < WIDTH; i++ {
		b.Field[i] = make([]int, HEIGHT)
	}
}

func (l *Layout) AllocateNetwork(nl *LayoutNetwork, kind string) {
	switch kind {
	case "actor":
		l.Actor = append(l.Actor, nl)
	case "main":
		l.Main = append(l.Main, nl)
	case "others":
		l.Others = append(l.Others, nl)
	}
}

func (l *Layout) CalcNetworksPostion() {
	l.setActorWidth()
	l.setMainWidth()
	l.calcActorPostion()
	l.calcMainPostion()
	l.calcOthersPostion()
}

func (l *Layout) calcActorPostion() {
	posX := LAYOUTMERGIN
	posY := LAYOUTMERGIN
	for _, v := range l.Actor {
		v.SetPosition(posX, posY)
		posY += v.Height + NETMARGIN
	}
}

func (l *Layout) calcMainPostion() {
	posX := LAYOUTMERGIN + l.ActorWidth + LAYOUTMERGIN
	posY := LAYOUTMERGIN
	for _, v := range l.Main {
		v.SetPosition(posX, posY)
		posY += v.Height + NETMARGIN
	}
}

func (l *Layout) calcOthersPostion() {
	posX := 3*LAYOUTMERGIN + l.ActorWidth + l.MainWidth
	posY := LAYOUTMERGIN
	for _, v := range l.Others {
		v.SetPosition(posX, posY)
		posY += v.Height + NETMARGIN
	}
}

func (l *Layout) setActorWidth() {
	for _, v := range l.Actor {
		if l.ActorWidth < v.Width {
			l.ActorWidth = v.Width
		}
	}
}

func (l *Layout) setMainWidth() {
	for _, v := range l.Main {
		if l.MainWidth < v.Width {
			l.MainWidth = v.Width
		}
	}
}

func (l *Layout) JoinNetworks() []*LayoutNetwork {
	nls := []*LayoutNetwork{}
	nls = append(nls, l.Actor...)
	nls = append(nls, l.Main...)
	nls = append(nls, l.Others...)
	return nls
}
