package config

import (
	"fmt"

	"github.com/CarlosBarbosaFilho/api-manager-projects/helper"
	"github.com/CarlosBarbosaFilho/api-manager-projects/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "manager_projects"
)

var (
	logger *Logger
)

func databaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbs, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorDefault(err)

	return dbs
}

func Init() {
	db := databaseConnection()
	err := db.AutoMigrate(&model.Project{})
	if err != nil {
		logger.Errorf("Error to create table %v", err)
	}
}
