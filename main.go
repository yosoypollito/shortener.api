package main

import(
	"github.com/gin-gonic/gin"
	"shortener/api/controllers/link"
)

func main(){

	//Configure database

	r := gin.Default()

	r.GET("/:id",link.GetById)
	r.POST("/link", link.CreateNew)
	r.Run()
}
