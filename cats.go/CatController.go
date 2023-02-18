package main

import (
 "github.com/gin-gonic/gin"
"log""
)

func GetAllCats(c *gin.Context) {
	c.JSON(200, FindAllCats())
}

func AddCat(c *gin.Context) {
	var dog Cat

	if err := c.BindJSON(&dog); err != nil {
		return
	}

	c.JSON(200, CreateCat(dog))
}

func GetCat(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, FindCatById(id))
}

func DeleteCat(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, DeleteCatById(id))
}

func EditCat(c *gin.Context) {
	id := c.Param("id")

	var cat Cat
	if err := c.BindJSON(&cat); err != nil {
		return
	}

	cat.ID = id

	log.Println(cat)

	c.JSON(200, UpdateCat(cat))
}

func 'UpdateCat'(cat cat) Cat {
	pgConnect := PostgresConnect()

	oldCat := FindCatById(cat.ID)

	oldCat.Name = cat.Name
	oldCat.IsStripe = cat.IsStripe
	oldCat.Color = cat.Color

	 err := pgConnect.Model(&oldCat).Select()

		Where(condition: "id = ?",  oldCat.ID).
		Update()
	if err != nil {
		panic(err)
	}
		}
