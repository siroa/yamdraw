package service

import (
	"generate/domain/layout"
	"generate/domain/yaml"
	"generate/utils/logger"
	"generate/utils/spliter"
	"strconv"
	"strings"
)

type DefaultLayoutService struct {
	Layout *layout.Layout
}

func (l DefaultLayoutService) CreateNetworkLayout(networks *[]yaml.Networks) []*layout.LayoutNetwork {
	nls := []*layout.LayoutNetwork{}
	g := layout.Layout{}
	for _, v := range *networks {
		id, name, kind, c := v.Network.ReadNetwork()
		netLayout := layout.CreateLayoutNetwork(id, name)
		if c == nil {
			netLayout.CalcSize()
			g.AllocateNetwork(netLayout, kind)
			continue
		}
		lServers := CreateServerLayout(c)
		netLayout.Servers = lServers
		netLayout.CalcSize()
		netLayout.CalcServerPosition()
		g.AllocateNetwork(netLayout, kind)
	}
	g.CalcNetworksPostion()
	nls = g.JoinNetworks()
	return nls
}

func (l DefaultLayoutService) CreateProcedure(pl map[string]string) []layout.LayoutProcedure {
	ps := []layout.LayoutProcedure{}

	for k, v := range pl {
		route := strings.ToLower(k)
		procedure := spliter.ProcedureSplit(v)
		// ToDo: entry pointを選ぶようにする
		rps := extractProcedure(route, procedure)
		ps = append(ps, rps...)
	}
	return ps
}

// a~xのルート情報とルートの順番が記載された文字配列から手順のレイアウトを決める
func extractProcedure(r string, p []string) []layout.LayoutProcedure {
	ps := []layout.LayoutProcedure{}
	num := len(p) / 2
	n := 0
	for i := 0; i < num; i++ {
		route := r + strconv.Itoa(i+1)
		// ToDo: entry pointを選ぶようにする
		lp, err := layout.CreateLayoutProcedure(route, p[n], p[n+2], p[n+1], 3, 2)
		if err != nil {
			// ToDo: エラーハンドリングを検討する
			logger.Error("Can't create struct. : " + err.Error())
		}
		ps = append(ps, lp)
		n += 2
	}
	return ps
}

func CreateServerLayout(c *[]yaml.Servers) []*layout.LayoutServer {
	lServers := []*layout.LayoutServer{}
	for _, v := range *c {
		id, name, mws := v.Server.ReadServer()
		ls := layout.CreateLayoutServer(id, name)
		if mws == nil {
			ls.CalcSize()
			lServers = append(lServers, ls)
			continue
		}
		lMws := CreateMws(mws)
		ls.Mws = lMws
		ls.CalcSize()
		ls.CalcMwPostion()
		lServers = append(lServers, ls)
	}
	return lServers
}

func CreateMws(mws *[]yaml.Mws) []*layout.LayoutMw {
	lMws := []*layout.LayoutMw{}
	for _, v := range *mws {
		id, name, accessories, err := v.Mw.ReadMw()
		if err != nil {
			continue
		}
		lMw := layout.CreateLayoutMw(id, name)
		lMw = createAccessories(accessories, lMw)
		lMw.CalcWidth()
		lMw.CalcHeight()
		lMw.CalcAccessoriesPosion()
		lMws = append(lMws, lMw)
	}
	return lMws
}

func createAccessories(acs map[string]*[]map[interface{}]interface{}, lMw *layout.LayoutMw) *layout.LayoutMw {
	if len(acs) == 0 {
		return lMw
	}
	for k, v := range acs {
		if v == nil {
			continue
		}
		for _, v := range *v {
			i, ok := v["id"].(string)
			if ok == false {
				num, _ := v["id"].(int)
				i = strconv.Itoa(num)
			}
			// ToDo: need error handling
			n, _ := v["name"].(string)
			a := layout.CreateLayoutAccessory(i, n, k).SetSize()
			lMw.Accessories = append(lMw.Accessories, a)
		}
	}
	return lMw
}
