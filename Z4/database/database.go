package database

import (
	"store_backend/environment"
	"store_backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize(env environment.Environment) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(env.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.Cart{},
	)

	if err != nil {
		panic(err)
	}

	if env.ENV == environment.Development {
		Seed(db)
	}

	return db
}
