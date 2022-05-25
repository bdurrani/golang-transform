package cleanup

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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

func StartCleanup() {
	go SessionCleanupTask()
}

func SessionCleanupTask() {
	quit := make(chan os.Signal, 1) // buffered
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		select {
		case sig := <-quit:
			log.Println("all done", sig)
			return
		case <-time.After(time.Second * 5):
			log.Println("peek: SessionCleanupTask")
		}
	}
}
