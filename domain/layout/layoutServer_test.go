package layout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcSize(t *testing.T) {
	// mwが0個の場合
	s := CreateLayoutServer("1", "server1")
	s.CalcSize()
	require.Equal(t, SEVERWIDTH, s.Width)
	require.Equal(t, SERVERHEIGHT, s.Height)

	// mwが1つの場合
	mw := CreateLayoutMw("mw1", "mw1")
	mwWid1 := 100
	mwHei1 := 120
	mw.SetSize(mwWid1, mwHei1)
	s.Mws = append(s.Mws, mw)
	s.CalcSize()
	want := mwWid1 + MWMARGIN
	require.Equal(t, want, s.Width)
	want = int(float64(mwHei1) * MWHEIGHTSCALE)
	require.Equal(t, want, s.Height)

	// mwが2個の場合
	mw2 := CreateLayoutMw("mw2", "mw2")
	mwWid2 := 100
	mwHei2 := 110
	mw2.SetSize(mwWid2, mwHei2)
	s.Mws = append(s.Mws, mw2)
	s.CalcSize()
	want = 2 * (mwWid2 + MWMARGIN)
	require.Equal(t, want, s.Width)
	want = int(float64(mwHei1) * MWHEIGHTSCALE)
	require.Equal(t, want, s.Height)
}
