package common

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"product-service/app/infra/db"
)

func BeginTx() *gorm.DB {
	return db.DB().Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)}).Begin()
}

func RecoveryTx(tx *gorm.DB) {
	if err := recover(); err != nil {
		tx.Rollback()
		panic(err)
	}
}
