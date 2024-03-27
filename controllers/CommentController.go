package controllers

import (
	"mygramapi/models"
	"mygramapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *commentController {
	return &commentController{commentService: commentService}
}

func (c *commentController) PostComment(ctx *gin.Context) {
	var result gin.H
	input := models.Comment{}
	data, ok := ctx.MustGet("userData").(map[string]interface{})
	if !ok {
		result = gin.H{
			"error": "Failed to get user data",
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	userId := data["id"].(float64)

	getUser, erri := c.commentService.FindUserById(int(userId))

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

	photo, err := c.commentService.Create(&getUser, input)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusCreated, photo)
}

func (c *commentController) GetComment(ctx *gin.Context) {
	var result gin.H
	photos, err := c.commentService.FindAll()

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

func (c *commentController) GetCommentById(ctx *gin.Context) {
	var result gin.H
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	comment, err := c.commentService.FindCommentById(id)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func (c *commentController) UpdateComment(ctx *gin.Context) {
	var (
		result     gin.H
		commentReq models.CommentUpdate
	)
	if err := ctx.ShouldBindJSON(&commentReq); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := c.commentService.Update(id, commentReq)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *commentController) DeleteComment(ctx *gin.Context) {
	var result gin.H

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := c.commentService.Delete(id)

	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	result = gin.H{
		"message": "Your comment has been successfully deleted",
	}
	ctx.JSON(http.StatusOK, result)
}
