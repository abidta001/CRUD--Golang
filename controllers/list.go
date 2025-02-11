package controllers

import (
	"admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPerson(c *gin.Context) {
	var person []models.Person
	models.DB.Find(&person)

	c.JSON(http.StatusOK, gin.H{"person": person})

}
