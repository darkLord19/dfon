package configparser

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	fname := "testdata/config.json"
	_, err := LoadConfig(fname)
	if err != nil {
		fmt.Println("Err:", err.Error())
		t.Fail()
	}
}
