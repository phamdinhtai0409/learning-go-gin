package router

import (
	"learning-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

type BlogRouter struct {
	blogCtrl *controllers.BlogController
}

func NewBlogRouter(blogCtrl *controllers.BlogController) *BlogRouter {
	return &BlogRouter{blogCtrl: blogCtrl}
}

func (ctrl *BlogRouter) SetBlogRouter(rg *gin.RouterGroup) {
	g := rg.Group("blogs")

	g.GET("get", ctrl.blogCtrl.Get)
	g.POST("post", ctrl.blogCtrl.Post)
}
