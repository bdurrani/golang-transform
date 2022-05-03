package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

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
	c.SaveUploadedFile(file, "./test.jpg")

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func Index(c *gin.Context) {
	index := gin.H{
		"title": "Transformation",
	}
	c.HTML(http.StatusOK, "index.tmpl", index)
}
