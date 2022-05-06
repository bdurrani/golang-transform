package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	setupTmpDirectory("upload")
	// openFile, err := file.Open()
	// if err != nil {
	// 	panic(err)
	// }
	// defer openFile.Close()

	// tmpFile, err := os.CreateTemp("/tmp/uploads", "upload.*.tmp")

	// if err != nil {
	// 	c.String(http.StatusInternalServerError, "Unable to create temp file")
	// 	return
	// }
	// defer tmpFile.Close()

	// io.Copy(tmpFile, openFile)
	// Upload the file to specific dst.
	err := c.SaveUploadedFile(file, "./test.jpg")
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to save uploaded file")
		return
	}
	file.Open()
	// c.Data(http.StatusOK, "image/png", )
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func Index(c *gin.Context) {
	index := gin.H{
		"title": "Transformation",
	}
	c.HTML(http.StatusOK, "index.tmpl", index)
}

func setupTmpDirectory(dirName string) (string, error) {
	tmpPath := filepath.Join("/tmp", dirName)
	err := os.MkdirAll(tmpPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	return tmpPath, nil
}
