package controllers

import (
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GettingAllPhotos godoc
// @Summary Get details
// @Description Get details of all photos
// @Tags Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} models.Photo "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Router /photos [get]
func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	photos := []models.Photo{}
	photoId := c.Param("photoId")

	if photoId != "" {
		err := db.Where("id = ?", photoId).Find(&photos).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}

		c.JSON(http.StatusOK, photos[0])
		return
	}

	err := db.Find(&photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, photos)
}

// AddingPhoto godoc
// @Summary Post details for a given Id
// @Description Post details of photo corresponding to the input Id
// @Tags Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @param helpers.PhotoInput body helpers.PhotoInput true "create photo"
// @Success 201	{object} models.Photo "Created"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// UpdatingPhoto godoc
// @Summary Update photo identified by the given Id
// @Description Update the photo corresponding to the input Id
// @Tags Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param photoId path int true "ID of the photo to be updated"
// @param helpers.PhotoInput body helpers.PhotoInput true "update photo"
// @Success 200	{object} models.Photo "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /photos/{photoId} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletingPhoto godoc
// @Summary Delete photo identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param photoId path int true "ID of the photo to be deleted"
// @Success 200 {object} helpers.DeleteResponse "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /photos/{photoId} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Photo with id %v has been deleted", photoId),
	})
}

// GettingPhoto godoc
// @Summary Get details for a given Id
// @Description Get details photo corresponding to the input Id
// @Tags Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param photoId path int true "ID of the photo"
// @Success 200 {object} models.Photo "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /photos/{photoId} [get]
func HelloPhoto(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
