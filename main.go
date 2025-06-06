package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umpprats/gin-sysacad/config"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bienvenidos a SysAcad!",
		})
	})

	return r
}

func main() {
	config.ConnectionDB()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
