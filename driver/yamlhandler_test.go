package driver

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYamlHandler(t *testing.T) {
	// 正常読み込み

	testFilePath := filepath.Clean(`..\test\test01.yaml`)
	y, err := NewYamlHandler(testFilePath)
	require.NoError(t, err)
	want, err := ioutil.ReadFile(testFilePath)
	require.NoError(t, err)
	require.Equal(t, want, y.File)

	// ファイルが存在しない
	testFilePath = filepath.Clean(`..\test\hogehoge.yaml`)
	y, err = NewYamlHandler(testFilePath)
	require.Error(t, err)
}
