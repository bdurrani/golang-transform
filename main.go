package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/bdurrani/golang-transform/cleanup"
	"html/template"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/bdurrani/golang-transform/controllers"
	"github.com/gin-contrib/static"
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

	router.Use(static.Serve("/assets/", static.LocalFile("/tmp/upload", false)))

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
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	cleanup.StartCleanup(ctx)
	router := setupEngine()
	_ = router.SetTrustedProxies([]string{"127.0.0.1"})
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	<-ctx.Done()
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
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
