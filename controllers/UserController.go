package controllers

import (
	"fmt"
	"mygramapi/helpers"
	"mygramapi/models"
	"mygramapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService: userService}
}

func (c *userController) RegisterUser(ctx *gin.Context) {
	var result gin.H
	input := models.User{}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	user, err := c.userService.Create(input)
	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *userController) LoginUser(ctx *gin.Context) {
	contentType := ctx.Request.Header.Get("Content-Type")

	var user models.UserRequest

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	getUser, err := c.userService.Login(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(getUser.Password), []byte(user.Password))
	fmt.Println(getUser.Password, user.Password)
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	token := helpers.GenerateToken(uint(getUser.ID), user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *userController) GetUsers(ctx *gin.Context) {
	var result gin.H
	users, err := c.userService.FindAll()

	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	result = gin.H{
		"data": users,
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *userController) GetUserById(ctx *gin.Context) {
	var result gin.H
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := c.userService.FindById(id)

	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	if user.Age == 0 {
		result = gin.H{
			"error": "user not found",
		}
		ctx.JSON(http.StatusNotFound, result)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var (
		result  gin.H
		userReq models.UserUpdate
	)
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		result = gin.H{
			"error": err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := c.userService.Update(id, userReq)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	var result gin.H

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := c.userService.Delete(id)

	if err != nil {
		result = gin.H{
			"error": err,
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	result = gin.H{
		"message": "Your account has been successfully deleted",
	}
	ctx.JSON(http.StatusOK, result)
}
