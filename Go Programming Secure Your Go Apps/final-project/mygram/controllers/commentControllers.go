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

// GettingAllComments godoc
// @Summary Get details
// @Description Get details of all comments
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Success 200 {object} models.Comment "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Router /comments [get]
func GetComment(c *gin.Context) {
	db := database.GetDB()
	comments := []models.Comment{}
	commentId := c.Param("commentId")

	if commentId != "" {
		err := db.Where("id = ?", commentId).Find(&comments).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}

		c.JSON(http.StatusOK, comments[0])
		return
	}

	err := db.Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, comments)
}

// AddingComment godoc
// @Summary Post details for a given Id
// @Description Post details of comment corresponding to the input Id
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @param helpers.CommentInput body helpers.CommentInput true "create comment"
// @Success 201	{object} models.Comment "Created"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// UpdatingComment godoc
// @Summary Update comment identified by the given Id
// @Description Update the comment corresponding to the input Id
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param commentId path int true "ID of the comment to be updated"
// @param helpers.CommentInput body helpers.CommentInput true "update comment"
// @Success 200	{object} models.Comment "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /comments/{commentId} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{PhotoID: Comment.PhotoID, Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeletingComment godoc
// @Summary Delete comment identified by the given id
// @Description Delete the order corresponding to the input id
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param commentId path int true "ID of the comment to be deleted"
// @Success 200 {object} helpers.DeleteResponse "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Comment with id %v has been deleted", commentId),
	})
}

// GettingComment godoc
// @Summary Get details for a given Id
// @Description Get details comment corresponding to the input Id
// @Tags Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Param commentId path int true "ID of the comment"
// @Success 200 {object} models.Comment "OK"
// @Failure	401	{object} helpers.APIError "Unauthorized"
// @Failure	400	{object} helpers.APIError "Bad Request"
// @Failure	500	{object} helpers.APIError "Server Error"
// @Router /comments/{commentId} [get]
func HelloComment(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
