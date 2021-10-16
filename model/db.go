package model

import (
	"GinBlog/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

var db *gorm.DB
var err error
func InitDB(){
	dbLink:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		)
	fmt.Println(dbLink)
	db,err=gorm.Open(mysql.Open(dbLink),&gorm.Config{
		//gorm 日志模式 slient
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务
		SkipDefaultTransaction: true,
	})
	if err!=nil {
		fmt.Println("打开数据库错误,", err)
		os.Exit(1)
	}
	_ =db.AutoMigrate(&User{},&Article{},&Category{})

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
