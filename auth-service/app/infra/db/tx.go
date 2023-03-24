package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func BeginTx() *gorm.DB {
	return DB().Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)}).Begin()
}

func RecoveryTx(tx *gorm.DB) {
	if err := recover(); err != nil {
		tx.Rollback()
		panic(err)
	}
}
