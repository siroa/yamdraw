package layout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLayoutMwWidth(t *testing.T) {
	// DBとprocessの数が0
	mw := CreateLayoutMw("test01", "testname1")
	mw.CalcWidth()
	require.Equal(t, DBWIDTH+MARGIN, mw.Width)

	// dbの数とprocessの数が同数
	db := CreateLayoutDB("db1", "dbname1")
	mw.DB = append(mw.DB, db)
	pr := CreateLayoutProcess("process1", "processname1")
	mw.Process = append(mw.Process, pr)
	mw.CalcWidth()
	require.Equal(t, PROWIDTH+MARGIN, mw.Width)

	// dbの個数がprocessより多い
	db2 := CreateLayoutDB("db1", "dbname1")
	mw.DB = append(mw.DB, db2)
	mw.CalcWidth()
	require.Equal(t, len(mw.DB)*(DBHEIGHT+MARGIN), mw.Width)
}

func TestLayoutMwHeight(t *testing.T) {
	// DBとprocessの数が0
	mw := CreateLayoutMw("test01", "testname1")
	mw.CalcHeight()
	require.Equal(t, int(MWHEIGHTSCALE*float64(DBHEIGHT+20)), mw.Height)
	// Process=1, db=0
	pr := CreateLayoutProcess("process1", "processname1")
	mw.Process = append(mw.Process, pr)
	mw.CalcHeight()
	require.Equal(t, int(MWHEIGHTSCALE*float64(DBHEIGHT+20)), mw.Height)

	// Process=1, db=1
	db := CreateLayoutDB("db1", "dbname1")
	mw.DB = append(mw.DB, db)
	mw.CalcHeight()
	sumHeight := DBHEIGHT + PROHEIGHT + MARGIN
	require.Equal(t, int(MWHEIGHTSCALE*float64(sumHeight)), mw.Height)
}

func TestCalcDBPosition(t *testing.T) {
	mw := CreateLayoutMw("test01", "testname1")
	db := CreateLayoutDB("db1", "dbname1")
	mw.DB = append(mw.DB, db)
	db2 := CreateLayoutDB("db2", "dbname2")
	mw.DB = append(mw.DB, db2)
	mwW := 100
	mwH := 140
	mw.Width = mwW
	mw.Height = mwH
	mw.CalcDBPostion()
	want := (mwW - DBWIDTH*len(mw.DB)) / (len(mw.DB) + 1)
	require.Equal(t, want, mw.DB[0].PositionX)
	want += want + DBWIDTH
	require.Equal(t, want, mw.DB[1].PositionX)
	want = mwH - (DBHEIGHT + 10)
	require.Equal(t, want, mw.DB[0].PositionY)
	require.Equal(t, want, mw.DB[1].PositionY)

	pr := CreateLayoutProcess("process1", "processname1")
	mw.Process = append(mw.Process, pr)
	mw.CalcDBPostion()
	want = mwH - (20 + DBHEIGHT + PROHEIGHT)
	require.Equal(t, want, mw.DB[0].PositionY)
	require.Equal(t, want, mw.DB[1].PositionY)
}

func TestCalcProcessPosition(t *testing.T) {
	mw := CreateLayoutMw("test01", "testname1")
	pr := CreateLayoutProcess("process1", "processname1")
	mw.Process = append(mw.Process, pr)
	pr2 := CreateLayoutProcess("process2", "processname2")
	mw.Process = append(mw.Process, pr2)
	mwW := 100
	mwH := 140
	mw.Width = mwW
	mw.Height = mwH
	mw.CalcProcessPostion()
	want := (mwW - PROWIDTH*len(mw.Process)) / (len(mw.Process) + 1)
	require.Equal(t, want, mw.Process[0].PositionX)
	want += want + PROWIDTH
	require.Equal(t, want, mw.Process[1].PositionX)
	want = mwH - (PROHEIGHT + 10)
	require.Equal(t, want, mw.Process[0].PositionY)
	require.Equal(t, want, mw.Process[1].PositionY)

	db := CreateLayoutDB("db1", "dbname1")
	mw.DB = append(mw.DB, db)
	mw.CalcProcessPostion()
	want = mwH - (10 + PROHEIGHT)
	require.Equal(t, want, mw.Process[0].PositionY)
	require.Equal(t, want, mw.Process[1].PositionY)
}
