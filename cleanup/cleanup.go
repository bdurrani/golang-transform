package cleanup

import (
	"fmt"
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
			info, err := file.Info()

			if err != nil {
				continue
			}

			fmt.Println("Last modified time : ", info.ModTime())
		}
	}
	return filesToDelete, nil
}
