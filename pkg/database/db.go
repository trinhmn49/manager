package database

import (
	"fmt"
	"manager/pkg/config"
	"manager/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	db *gorm.DB
}

func NewDb(config *config.DbConfig) *DbInstance {
	instance := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", config.Host, config.Port, config.UserName, config.Password, config.DbName)
	db, err := gorm.Open(postgres.Open(instance), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	logger.Info("Connected db success")
	return &DbInstance{
		db: db,
	}
}
