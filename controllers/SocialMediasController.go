package controllers

import (
	"mygramapi/models"
	"mygramapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	socialMediaService service.SocialMediasService
}

func NewSocialMediaController(socialMediaService service.SocialMediasService) *socialMediaController {
	return &socialMediaController{socialMediaService: socialMediaService}
}

func (c *socialMediaController) PostSocialMedia(ctx *gin.Context) {
	var result gin.H
	input := models.SocialMedia{}
	data, ok := ctx.MustGet("userData").(map[string]interface{})
	if !ok {
		result = gin.H{
			"error": "Failed to get user data",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	userId := data["id"].(float64)

	getUser, erri := c.socialMediaService.FindUserById(int(userId))

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

	socialMedia, err := c.socialMediaService.Create(&getUser, input)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusCreated, socialMedia)
}

func (c *socialMediaController) GetSocialMedia(ctx *gin.Context) {
	var result gin.H
	socialMedias, err := c.socialMediaService.FindAll()

	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	result = gin.H{
		"data": socialMedias,
	}
	ctx.JSON(http.StatusOK, socialMedias)
}

func (c *socialMediaController) GetSocialMediaById(ctx *gin.Context) {
	var result gin.H
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	sm, err := c.socialMediaService.FindSocialMediaById(id)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": sm,
	})
}

func (c *socialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	var (
		result gin.H
		smReq  models.SocialMedia
	)
	if err := ctx.ShouldBindJSON(&smReq); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	sm, err := c.socialMediaService.Update(id, smReq)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusOK, sm)
}

func (c *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	var result gin.H

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := c.socialMediaService.Delete(id)

	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	result = gin.H{
		"message": "Your social media has been successfully deleted",
	}
	ctx.JSON(http.StatusOK, result)
}
