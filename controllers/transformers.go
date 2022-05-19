package controllers

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func HandleUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	mimeType := mime.TypeByExtension(filepath.Ext(file.Filename))
	log.Println(file.Filename + " " + mimeType)

	tmpDirectory, err := setupTmpDirectory("upload")
	if err != nil {
		return
	}

	id := uuid.New()
	// Upload the file to specific dst.
	saveLocation := filepath.Join(tmpDirectory, id.String())
	err = c.SaveUploadedFile(file, saveLocation)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to save uploaded file")
		return
	}
	// data, err := ioutil.ReadFile(saveLocation)
	// if err != nil {
	// 	return
	// }

	redirUrl := fmt.Sprintf("/assets/%s", id.String())
	c.Redirect(http.StatusMovedPermanently, redirUrl)
	// c.Data(http.StatusOK, mimeType, data)
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
