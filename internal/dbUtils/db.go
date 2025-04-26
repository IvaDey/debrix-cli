package dbUtils

import (
	"github.com/glebarez/sqlite"
	"github.com/ivadey/debrix-cli/internal/utils"
	"gorm.io/gorm"
)

var connection *gorm.DB

func OpenDb() *gorm.DB {
	if connection != nil {
		return connection
	}

	config := utils.GetConfig()

	var err error
	connection, err = gorm.Open(sqlite.Open(config.DbPath), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err = connection.AutoMigrate(&StoredTodo{}); err != nil {
		panic("failed to migrate database")
	}

	return connection
}
