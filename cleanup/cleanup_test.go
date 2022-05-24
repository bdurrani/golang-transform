package cleanup

import (
	"testing"
)

func TestCleanup(t *testing.T) {
	t.Run("clean up works", func(t *testing.T) {
		filesToCleanup("/tmp")
	})
}
