package main

import (
	"github.com/bdurrani/golang-transform/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

var db = map[string]string{
	"test": "test-value",
}

func setupEngine() *gin.Engine {
	gin.DisableConsoleColor()
	router := gin.Default()

	control := controllers.Controller{
		Db: db,
	}
	router.GET("/user/:name", control.GetUser)
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
