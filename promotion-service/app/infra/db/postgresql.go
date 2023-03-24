package db

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"

	"promotion-service/app/config"
)

var postgresDB *gorm.DB

func InitPostgres() {
	conf := config.Config
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Postgres.UserName,
		conf.Postgres.Password,
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.Database,
	)

	var err error
	postgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
func DB() *gorm.DB {
	if postgresDB == nil {
		InitPostgres()
	}
	return postgresDB
}
func ClosePostgres() {
	if db, _ := postgresDB.DB(); db != nil {
		if err := db.Close(); err != nil {
			fmt.Println("[ERROR] Cannot close mysql connection, err:", err)
		}
	}
}
