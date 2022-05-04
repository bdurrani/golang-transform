package controllers

import (
	"github.com/google/uuid"
	"os"
	"testing"
)

func TestTransformers(t *testing.T) {
	t.Run("setting up temp directory succeeds", func(t *testing.T) {
		tmpDirName := uuid.New().String()
		err := setupTmpDirectory(tmpDirName)
		if err != nil {
			t.Errorf("err was not null")
		}
		_, err = os.Stat("/tmp/" + tmpDirName)
		exists := !os.IsNotExist(err)
		if exists != true {
			t.Errorf("got %t, wanted %t", exists, true)
		}
		_ = os.Remove("/tmp" + tmpDirName)
	})
}
