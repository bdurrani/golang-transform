package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NameValueMap map[string]string
type Controller struct {
	Db NameValueMap
}

func (base *Controller) GetUser(context *gin.Context) {
	context.String(http.StatusOK, "pong")
	fmt.Printf("ClientIP: %s\n", context.ClientIP())
	user := context.Params.ByName("name")
	value, ok := base.Db[user]
	if ok {
		context.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		context.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}
