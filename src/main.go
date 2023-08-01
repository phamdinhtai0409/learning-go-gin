package main

import (
	"fmt"
	"learning-go-gin/config"
	"learning-go-gin/database"
	"learning-go-gin/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	conf := config.NewConfig()
	fmt.Println(conf.DB.Host)

	dbConn := database.NewConnection(conf)

	r := gin.Default()

	app := initializeApp(conf, dbConn, r)
	app.start(conf)
}

func NewApp(conf config.Config, r *router.Router) App {
	return App{
		conf: conf,
		r:    r,
	}
}

type App struct {
	conf config.Config
	r    *router.Router
}

func (app *App) start(conf config.Config) {
	app.r.Setup(conf)
	app.r.Run(fmt.Sprintf(":%d", app.conf.App.Port))
}
