package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	var connection string = "host=localhost port=5432 user=root password=root dbname=assignmentfinal sslmode=disable"
	return gorm.Open(postgres.Open(connection), &gorm.Config{})
}
