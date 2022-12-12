package spliter

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcedureSplit(t *testing.T) {
	str := "1->2.>3"
	want := []string{"1", "->", "2", ".>", "3"}
	ss := ProcedureSplit(str)
	require.Equal(t, want, ss)
}

func TestSp(t *testing.T) {
	re := regexp.MustCompile(`->|.>`)
	var ss []string
	ss = nil
	require.Equal(t, ss, Split("", 0, re))
	s := ""
	require.Equal(t, []string{""}, Split(s, -1, re))
}
