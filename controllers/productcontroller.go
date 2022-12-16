package controllers

import (
	"net/http"
	"service-go-restapi/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var staff []models.Staff

	models.DB.Find(&staff)
	c.JSON(http.StatusOK, gin.H{
		"Data Karyawan": staff,
	})
}
