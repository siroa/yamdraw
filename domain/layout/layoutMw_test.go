package layout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLayoutMwWidth(t *testing.T) {
	// Number of DBs and processes is 0
	mw := CreateLayoutMw("test01", "testname1")
	mw.CalcWidth()
	require.Equal(t, DBWIDTH+MARGIN, mw.Width)

	// Equal number of db's and process
	db := CreateLayoutAccessory("db1", "dbname1", DBWIDTH, DBWIDTH)
	mw.DB = append(mw.DB, db)
	pr := CreateLayoutAccessory("process1", "processname1", PROWIDTH, PROHEIGHT)
	mw.Process = append(mw.Process, pr)
	mw.CalcWidth()
	require.Equal(t, PROWIDTH+MARGIN, mw.Width)

	// Number of db's is more than process.
	db2 := CreateLayoutAccessory("db1", "dbname1", DBWIDTH, DBWIDTH)
	mw.DB = append(mw.DB, db2)
	mw.CalcWidth()
	require.Equal(t, len(mw.DB)*(DBHEIGHT+MARGIN), mw.Width)
}

func TestLayoutMwHeight(t *testing.T) {
	// Number of DBs and processes is 0
	mw := CreateLayoutMw("test01", "testname1")
	mw.CalcHeight()
	require.Equal(t, int(MWHEIGHTSCALE*float64(DBHEIGHT+20)), mw.Height)
	// Process=1, db=0
	pr := CreateLayoutAccessory("process1", "processname1", PROWIDTH, PROHEIGHT)
	mw.Process = append(mw.Process, pr)
	mw.CalcHeight()
	require.Equal(t, int(MWHEIGHTSCALE*float64(DBHEIGHT+20)), mw.Height)

	// Process=1, db=1
	db := CreateLayoutAccessory("db1", "dbname1", DBWIDTH, DBWIDTH)
	mw.DB = append(mw.DB, db)
	mw.CalcHeight()
	sumHeight := DBHEIGHT + PROHEIGHT + MARGIN
	require.Equal(t, int(MWHEIGHTSCALE*float64(sumHeight)), mw.Height)
}

func TestCalcDBPosition(t *testing.T) {
	mw := CreateLayoutMw("test01", "testname1")
	db := CreateLayoutAccessory("db1", "dbname1", DBWIDTH, DBWIDTH)
	mw.DB = append(mw.DB, db)
	db2 := CreateLayoutAccessory("db2", "dbname2", DBWIDTH, DBWIDTH)
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

	pr := CreateLayoutAccessory("process1", "processname1", PROWIDTH, PROHEIGHT)
	mw.Process = append(mw.Process, pr)
	mw.CalcDBPostion()
	want = mwH - (20 + DBHEIGHT + PROHEIGHT)
	require.Equal(t, want, mw.DB[0].PositionY)
	require.Equal(t, want, mw.DB[1].PositionY)
}

func TestCalcProcessPosition(t *testing.T) {
	mw := CreateLayoutMw("test01", "testname1")
	pr := CreateLayoutAccessory("process1", "processname1", PROWIDTH, PROHEIGHT)
	mw.Process = append(mw.Process, pr)
	pr2 := CreateLayoutAccessory("process2", "processname2", PROWIDTH, PROHEIGHT)
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

	db := CreateLayoutAccessory("db1", "dbname1", DBWIDTH, DBWIDTH)
	mw.DB = append(mw.DB, db)
	mw.CalcProcessPostion()
	want = mwH - (10 + PROHEIGHT)
	require.Equal(t, want, mw.Process[0].PositionY)
	require.Equal(t, want, mw.Process[1].PositionY)
}
