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

func StartCleanup(quit chan os.Signal) {
	go sessionCleanupTask1(quit)
}

func SessionCleanupTask() {
	quit := make(chan os.Signal, 1) // buffered
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		select {
		case sig := <-quit:
			log.Println("shut down clean up routine", sig)
			return
		case <-time.After(time.Second * 5):
			log.Println("peek: SessionCleanupTask")
		}
	}
}

func sessionCleanupTask1(quit chan os.Signal) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case sig := <-quit:
			log.Println("shut down clean up routine", sig)
			return
		case <-ticker.C:
			log.Println("peek: clean up tick")
		}
	}
}
