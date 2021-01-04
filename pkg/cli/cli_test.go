package cli

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestBleble(t *testing.T) {

	dummyStorage := &DummyStorage{
		ReadFN:  func(path string) ([]byte, error) { return nil, nil },
		WriteFN: func(path string, data []byte) error { return nil },
	}

	Tool := &CoreTool{Storage: dummyStorage}

	data, _ := Tool.List()

	assert.Equal(t, len(data), 0, "FAIL")

}

func TestBleble2(t *testing.T) {

	const testValue = "asdasd"

	Tool := &CoreTool{Storage: &DummyStorage{
		ReadFN: func(path string) ([]byte, error) { return nil, nil },
		WriteFN: func(path string, data []byte) error {
			stringData := string(data)
			assert.Equal(t, testValue, stringData, "cambio el formato")
			return nil
		},
	}}

	err := Tool.Save(testValue)

	assert.Nil(t, err, "FAIL")
}
