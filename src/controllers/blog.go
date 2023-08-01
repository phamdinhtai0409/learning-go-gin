package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct{}

func NewBlogController() *BlogController {
	return &BlogController{}
}

func (ctrl *BlogController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "blog get"})
}

func (ctrl *BlogController) Post(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "blog post"})
}
