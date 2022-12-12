package layout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNetworkCalcSize(t *testing.T) {
	// Zero number of servers
	n := CreateLayoutNetwork("network1", "network1")
	n.CalcSize()
	require.Equal(t, NETWIDTH, n.Width)
	require.Equal(t, NETHEIGHT, n.Height)
	// Number of servers 1
	s1 := CreateLayoutServer("server1", "server1")
	s1.Width = SEVERWIDTH
	s1.Height = SERVERHEIGHT
	n.Servers = append(n.Servers, s1)
	n.CalcSize()
	require.Equal(t, SEVERWIDTH+NETMARGIN, n.Width)
	require.Equal(t, SERVERHEIGHT+NETMARGIN*2, n.Height)
	// The number of servers is 2 and the later server is smaller
	s2 := CreateLayoutServer("server2", "server2")
	s2.Width = SEVERWIDTH - 20
	s2.Height = SERVERHEIGHT
	n.Servers = append(n.Servers, s2)
	n.CalcSize()
	require.Equal(t, SEVERWIDTH+NETMARGIN, n.Width)
	require.Equal(t, 2*(SERVERHEIGHT+NETMARGIN*2), n.Height)
}

func TestCalcServerPosition(t *testing.T) {
	n := CreateLayoutNetwork("network1", "network1")
	s1 := CreateLayoutServer("server1", "server1")
	s1.Width = SEVERWIDTH
	s1.Height = SERVERHEIGHT
	n.Servers = append(n.Servers, s1)
	n.CalcSize()
	n.CalcServerPosition()
	require.Equal(t, NETMARGIN/2, n.Servers[0].PositionX)
	require.Equal(t, n.Height-((NETMARGIN/2)+SERVERHEIGHT), n.Servers[0].PositionY)

}
