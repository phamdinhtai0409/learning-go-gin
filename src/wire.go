//go:build wireinject
// +build wireinject

package main

import (
	"learning-go-gin/config"
	"learning-go-gin/controllers"
	"learning-go-gin/router"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var superSet = wire.NewSet(
	controllers.Set,
	router.Set,
	NewApp,
)

func initializeApp(conf config.Config, dbConn *gorm.DB, r *gin.Engine) App {
	wire.Build(superSet)
	return App{}
}
