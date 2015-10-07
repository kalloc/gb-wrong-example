package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var engine *gin.Engine = gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.GET("/:name", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello %s\n", c.Param("name"))
	})
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Index\n")
	})
	engine.Run(":8080")
}
