package controllers

import (
	"mygramapi/models"
	"mygramapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *photoController {
	return &photoController{photoService: photoService}
}

func (c *photoController) PostPhoto(ctx *gin.Context) {
	var result gin.H
	input := models.Photo{}
	data, ok := ctx.MustGet("userData").(map[string]interface{})
	if !ok {
		result = gin.H{
			"error": "Failed to get user data",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	userId := data["id"].(float64)

	getUser, erri := c.photoService.FindUserById(int(userId))

	if erri != nil {
		result = gin.H{
			"error": erri.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	photo, err := c.photoService.Create(&getUser, input)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusCreated, photo)
}

func (c *photoController) GetPhotos(ctx *gin.Context) {
	var result gin.H
	photos, err := c.photoService.FindAll()

	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	result = gin.H{
		"data": photos,
	}
	ctx.JSON(http.StatusOK, photos)
}

func (c *photoController) GetPhotoById(ctx *gin.Context) {
	var result gin.H
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	photo, err := c.photoService.FindPhotoById(id)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}

func (c *photoController) UpdatePhoto(ctx *gin.Context) {
	var (
		result   gin.H
		photoReq models.PhotoUpdate
	)
	if err := ctx.ShouldBindJSON(&photoReq); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := c.photoService.Update(id, photoReq)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *photoController) DeletePhoto(ctx *gin.Context) {
	var result gin.H

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := c.photoService.Delete(id)

	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	result = gin.H{
		"message": "Your photo has been successfully deleted",
	}
	ctx.JSON(http.StatusOK, result)
}
