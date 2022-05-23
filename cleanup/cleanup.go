package cleanup

import (
	"os"
)

func filesToCleanup(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filesToDelete []string

	for _, file := range files {
		if !file.IsDir() {
			filesToDelete = append(filesToDelete, file.Name())
		}
	}
	return filesToDelete, nil
}
