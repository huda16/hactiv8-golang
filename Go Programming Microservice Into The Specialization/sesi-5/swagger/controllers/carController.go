package controllers

import (
	"fmt"
	"net/http"
	"swagger/database"
	"swagger/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// GetAllCars godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(c *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car datas :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// GetOneCars godoc
// @Summary Get details for a given Id
// @Description Get details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{id} [get]
func GetOneCars(c *gin.Context) {
	var db = database.GetDB()

	var carOne models.Car
	// err := db.Table("Car").Where("Id = ?", c.Param("id")).First(&car).Error
	err := db.First(&carOne, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": carOne})
}

// CreateCars godoc
// @Summary Post details for a given Id
// @Description Post details of car corresponding ti the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCars(c *gin.Context) {
	var db = database.GetDB()
	// validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create car
	carinput := models.Car{Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars, Pemilik: input.Pemilik}
	db.Create(&carinput)

	c.JSON(http.StatusOK, gin.H{"data": carinput})
}

// UpdateCars godoc
// @Summary Update car identified by the given Id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Param models.Car body models.Car true "update car"
// @Success 200 {object} models.Car
// @Router /cars/{id} [patch]
func UpdateCars(c *gin.Context) {
	var db = database.GetDB()

	var car models.Car
	err := db.First(&car, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create car
	carinput := models.Car{Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars, Pemilik: input.Pemilik}
	db.Where("id = ?", c.Param("id")).Updates(&carinput)

	c.JSON(http.StatusOK, gin.H{"data": carinput})
}

// DeletingCars godoc
// @Summary Delete car identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{id} [delete]
func DeleteCars(c *gin.Context) {
	var db = database.GetDB()
	// Get model if exist
	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
