package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"first_api/repository"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/assets", "./assets")

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{})
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{})
	})

	router.POST("/register", func(c *gin.Context) {
		name := c.PostForm("name")
		repository.InsertValue(name)
		ary := repository.GetAllPerson()
		c.HTML(http.StatusOK, "summary.tmpl", gin.H{
			"data": ary,
		})
	})

	router.GET("/summary", func(c *gin.Context) {
		ary := repository.GetAllPerson()
		c.HTML(http.StatusOK, "summary.tmpl", gin.H{
			"data": ary,
		})
	})

	router.POST("/summary/:id", func(c *gin.Context) {
		id := c.Param("id")
		repository.DeletePerson(id)
		ary := repository.GetAllPerson()
		c.HTML(http.StatusOK, "summary.tmpl", gin.H{
			"data": ary,
		})
	})

	router.Run()
}
