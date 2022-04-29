package main

import (
	"fmt"
	"github.com/bdurrani/golang-transform/controllers"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

var db = map[string]string{
	"test": "test-value",
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func setupEngine() *gin.Engine {
	gin.DisableConsoleColor()
	router := gin.Default()

	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./templates/raw.tmpl")
	control := controllers.Controller{
		Db: db,
	}
	router.GET("/user/:name", control.GetUser)
	router.POST("/upload", controllers.HandleUpload)

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl",
			gin.H{
				"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
			})
	})

	return router
}

func main() {
	config, err := GetConfig()
	checkerr(err)
	router := setupEngine()

	_ = router.SetTrustedProxies([]string{"127.0.0.1"})
	_ = router.Run(":" + config.Port)
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
