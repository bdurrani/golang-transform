package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	tmpDirectory, err := setupTmpDirectory("upload")
	if err != nil {
		return
	}

	// Upload the file to specific dst.
	saveLocation := filepath.Join(tmpDirectory, file.Filename)
	err = c.SaveUploadedFile(file, saveLocation)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to save uploaded file")
		return
	}
	data, err := ioutil.ReadFile(saveLocation)
	if err != nil {
		return
	}

	c.Data(http.StatusOK, "image/png", data)
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
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
