package driver

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDatabase() (db *gorm.DB) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		"postgres",
		"root",
		"root",
		"localhost",
		5432,
		"rust-database",
		"disable",
	)

	dialector := postgres.Open(dsn)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("db connection failed")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	db = db.Debug()

	return
}
