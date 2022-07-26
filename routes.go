package main

import (
	"github.com/gin-gonic/gin"
)

func (app *jsPool) routes() *gin.Engine {
	router := gin.Default()
	router.POST("/git", app.GitHandler)
	return router
}
