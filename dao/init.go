package dao

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  //string类型的字段长度
		DisableDatetimePrecision:  true, //
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	//主从配置
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connWrite)}, //读操作
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)},
			Policy:   dbresolver.RandomPolicy{},
		}))
	migration()

}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
