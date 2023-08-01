package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (ctrl *UserController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "user get"})
}

func (ctrl *UserController) Post(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "user post"})
}
