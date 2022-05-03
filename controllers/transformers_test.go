package controllers

import (
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("hello nurse", func(t *testing.T) {
		setupTmpDirectory("upload")
		_, err := os.Stat("/path/to/whatever")
		exists := !os.IsNotExist(err)
		want := true
		if exists != want {
			t.Errorf("got %t, wanted %t", exists, want)
		}
	})
}
