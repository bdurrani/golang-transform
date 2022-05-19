package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/bdurrani/golang-transform/controllers"
	"github.com/gin-gonic/gin"
)

var db = map[string]string{
	"test": "test-value",
}

//go:embed templates/*
var f embed.FS

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func setupEngine() *gin.Engine {
	gin.DisableConsoleColor()
	router := gin.Default()

	templates, err := loadTemplate()
	checkerr(err)

	router.SetHTMLTemplate(templates)
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLGlob("templates/*.tmpl")
	control := controllers.Controller{
		Db: db,
	}

	router.StaticFS("assets", http.Dir("/tmp/upload"))

	router.GET("/user/:name", control.GetUser)
	router.GET("/", controllers.Index)
	router.POST("/upload", controllers.HandleUpload)

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl",
			gin.H{
				"now": time.Now(),
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

func loadTemplate() (*template.Template, error) {
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	return templ, nil
	// t := template.New("")
	// for name, file := range Assets.Files {
	// 	if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
	// 		continue
	// 	}
	// 	h, err := ioutil.ReadAll(file)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	t, err = t.New(name).Parse(string(h))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// return t, nil
}
