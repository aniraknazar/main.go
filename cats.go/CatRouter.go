package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10/types"
	"log"
)

func GetAllCats(c *gin.Context ) {
	c.JSON(200, FindAllCats())
}
func AddCat(c *gin.Context) {
	var cat Cat

	if err := c.BindJSON(scat); err != nil:

		c.JSON(200,)
}
