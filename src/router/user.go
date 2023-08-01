package router

import (
	"learning-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userCtrl *controllers.UserController
}

func NewUserRouter(userCtrl *controllers.UserController) *UserRouter {
	return &UserRouter{userCtrl: userCtrl}
}

func (ctrl *UserRouter) SetUserRouter(rg *gin.RouterGroup) {
	g := rg.Group("users")

	g.GET("get", ctrl.userCtrl.Get)
	g.POST("post", ctrl.userCtrl.Post)
}
