package link

import (
	"github.com/gin-gonic/gin"
	"shortener/api/structs/links"
	"net/http"
	"shortener/api/helpers"
)

func GetById(c *gin.Context){

	id := c.Param("id")

	links, err := links.Get(id)

	if err != nil {
		panic(err)
	}


	c.Redirect(http.StatusMovedPermanently, string(links[0].Scope))
}

func CreateNew(c *gin.Context){
	var linkReq links.Body

	if err := c.BindJSON(&linkReq); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": helpers.GetErrors(err),
		})
		return
	}


	key, err := links.Create(linkReq.Scope)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"key":key,
	})
}
