package web

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func RunServer() {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome",
			"body":  "hello page",
		})
	})

	router.Run()
}
