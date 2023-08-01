package router

import (
	"learning-go-gin/config"
	"log"

	_ "learning-go-gin/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Set = wire.NewSet(
	NewRouter,
	NewUserRouter,
	NewBlogRouter,
)

type Router struct {
	router     *gin.Engine
	userRouter *UserRouter
	blogRouter *BlogRouter
}

func NewRouter(
	r *gin.Engine,
	userRouter *UserRouter,
	blogRouter *BlogRouter,
) *Router {
	return &Router{
		router:     r,
		userRouter: userRouter,
		blogRouter: blogRouter,
	}
}

const BasePath = "/api/v1"

//	@title			API Go Gin
//	@version		1.0
//	@description	API in Go using Gin framework.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:5050
// @BasePath	/api/v1
func (r *Router) Setup(conf config.Config) {
	router := r.router

	rg := router.Group(BasePath)

	r.userRouter.SetUserRouter(rg)
	r.blogRouter.SetBlogRouter(rg)

	rg.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (r *Router) Run(s string) {
	err := r.router.Run(s)
	if err != nil {
		log.Fatal(err)
	}
}
