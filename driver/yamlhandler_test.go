package driver

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYamlHandler(t *testing.T) {
	// normal mode read

	testFilePath := filepath.Clean(`..\test\test01.yaml`)
	y, err := NewYamlHandler(testFilePath)
	require.NoError(t, err)
	want, err := ioutil.ReadFile(testFilePath)
	require.NoError(t, err)
	require.Equal(t, want, y.File)

	// File does not exist
	testFilePath = filepath.Clean(`..\test\hogehoge.yaml`)
	y, err = NewYamlHandler(testFilePath)
	require.Error(t, err)
}
