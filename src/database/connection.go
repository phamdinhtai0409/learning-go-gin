package database

import (
	"fmt"
	"learning-go-gin/config"
	"learning-go-gin/model/entity"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection(conf config.Config) *gorm.DB {
	dbConf := conf.DB

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	logger := initLogger(conf)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		panic(fmt.Errorf("fatal error: %s\n", err))
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(fmt.Errorf("fatal error: %s\n", err))
	}

	return db
}

func initLogger(conf config.Config) logger.Interface {
	return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
}
