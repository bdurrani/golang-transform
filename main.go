package main

import (
	"github.com/bdurrani/golang-transform/controllers"
	"github.com/gin-gonic/gin"
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
	router := setupEngine()

	_ = router.SetTrustedProxies([]string{"127.0.0.1"})
	// Listen and Server in 0.0.0.0:8080
	_ = router.Run(":8080")
}
