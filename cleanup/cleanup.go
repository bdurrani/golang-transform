package cleanup

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
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

func StartCleanup(quit context.Context) {
	go sessionCleanupTask(quit)
}

func sessionCleanupTask(quit context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-quit.Done():
			log.Println("shut down clean up routine")
			return
		case <-ticker.C:
			log.Println("peek: clean up tick")
		}
	}
}
