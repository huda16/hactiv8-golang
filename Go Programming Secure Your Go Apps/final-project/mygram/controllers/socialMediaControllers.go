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

// GettingAllSocialMedia godoc
// @Summary Get details
// @Description Get details of all social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} models.SocialMedia "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Router /socialmedia [get]
func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMedia := []models.SocialMedia{}
	socialmediaId := c.Param("socialmediaId")

	if socialmediaId != "" {
		err := db.Where("id = ?", socialmediaId).Find(&socialMedia).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}

		c.JSON(http.StatusOK, socialMedia[0])
		return
	}

	err := db.Find(&socialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, socialMedia)
}

// AddingSocialMedia godoc
// @Summary Post details for a given Id
// @Description Post details of social media corresponding to the input Id
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @param helpers.SocialMediaInput body helpers.SocialMediaInput true "create social media"
// @Success 201	{object} models.SocialMedia "Created"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /socialmedia [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

// UpdatingSocialMedia godoc
// @Summary Update social media identified by the given Id
// @Description Update the social media corresponding to the input Id
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param socialmediaId path int true "ID of the social media to be updated"
// @param helpers.SocialMediaInput body helpers.SocialMediaInput true "update social media"
// @Success 200	{object} models.SocialMedia "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /socialmedia/{socialmediaId} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialmediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// DeletingSocialMedia godoc
// @Summary Delete social media identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param socialmediaId path int true "ID of the social media to be deleted"
// @Success 200 {object} helpers.DeleteResponse "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /socialmedia/{socialmediaId} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	SocialMedia := models.SocialMedia{}

	err := db.Where("id = ?", socialmediaId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Social Media with id %v has been deleted", socialmediaId),
	})
}

// GettingSocialMedia godoc
// @Summary Get details for a given Id
// @Description Get details social media corresponding to the input Id
// @Tags Social Media
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param socialmediaId path int true "ID of the social media"
// @Success 200 {object} models.SocialMedia "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /socialmedia/{socialmediaId} [get]
func HelloSocialMedia(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
