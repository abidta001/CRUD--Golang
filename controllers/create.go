package controllers

import (
	"admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPerson struct {
	Name      string `json:"name" binding:"required"`
	Age       int    `json:"age" binding:"required"`
	IsMarried bool   `json:"married" `
}

func CreatePerson(c *gin.Context) {
	var input createPerson
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person := models.Person{
		Name: input.Name, Age: input.Age, IsMarried: input.IsMarried,
	}
	models.DB.Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": person})

}
