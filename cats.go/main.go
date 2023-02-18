package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/dog/add", AddCat)

	r.GET("/api/dogs", GetAllCats)

	r.GET("/api/dog/:id", GetCat)

	r.DELETE("/api/dog/:id", DeleteCat)

	r.PUT("/api/dog/:id", EditCat)

	r.Run("0.0.0.0:8888")
}
