package main 

import (
	"github.com/git-gonic/gin"
	"github.com/google/uuid"
)

type Cat struct {
	ID string `json:"id"`
	Name string `json:"name"`
	IsStrip bool `json:"is_strip"`
	Color string `json:"color"`
}

func main () {
	var cats []Cat
    r := git.Default()
     
	r.GET("/api/cats", func(c *git.Context){
		cat := Cat{
			ID: unid.NewString(),
	        Name: "Барсик",
	        IsStrip: true,
	        Color: "orange",
		}
		cats = append(cats, cat)
		c.JSON(200, cat)
	})

	r.Run("0.0.0.8888")
}

