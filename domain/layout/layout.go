package layout

const (
	WIDTH  = 1000
	HEIGHT = 1000
)

type Layout struct {
	Ground      [][]int
	Actor       []*LayoutNetwork
	Main        []*LayoutNetwork
	Others      []*LayoutNetwork
	ActorWidth  int
	MainWidth   int
	OthersWidth int
}

type LayoutContents struct {
	ID        string
	Width     int
	Height    int
	PositionX int
	PositionY int
}

// 1 = 10px
// initial value is 0
func (l *Layout) InitGround() {
	l.Ground = make([][]int, WIDTH)
	for i := 0; i < WIDTH; i++ {
		l.Ground[i] = make([]int, HEIGHT)
	}
}

func (l *Layout) AllocateLayout(nls []*LayoutNetwork) {

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
	posX := 50
	posY := 50
	for _, v := range l.Actor {
		v.SetPosition(posX, posY)
		posY += v.Height + NETMARGIN
	}
}

func (l *Layout) calcMainPostion() {
	posX := 50 + l.ActorWidth + 50
	posY := 50
	for _, v := range l.Main {
		v.SetPosition(posX, posY)
		posY += v.Height + NETMARGIN
	}
}

func (l *Layout) calcOthersPostion() {
	posX := 50 + l.ActorWidth + 50 + l.MainWidth + 50
	posY := 50
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
