package config

import (
	"fast_demo/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db (GORM)：是面向业务的。你写代码时，用它来负责数据的 CRUD（增删改查）。
// 它让你不需要写复杂的 SQL 语句，直接用 Go 的结构体就能开开心心地跟数据打交道。

// sqlDB (Go 标准库)：是面向网络/底层物理通道的。它负责在幕后把控连接池的各项指标
// （比如最大能同时开多少条通道、闲着的时候保留几条通道、一条通道活多久等）。

func initDB() {
	dsn := AppConfig.Database.Dsn //获取数据库的信息
	//拿着这串 DSN 密码，让 GORM 框架去敲 MySQL 的大门
	//成功，则db是 GORM 封装好的高层级操作对象

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initial;ize database, got error: %v", err)
	}
	//让 GORM 把藏在它身后的 Go 官方“原生连接管理器”交出来
	//好让你去设置最大连接数等底层网络参数。
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to configure database, got error %v", err)
	}

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns) // 空闲时保留11条连接，避免频繁创建新连接
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenCons)  // 最多同时打开114条连接，超过的请求会排队等待
	sqlDB.SetConnMaxLifetime(time.Hour)                    // 每条连接最多活 1 小时，到期自动关闭换新的，防止连接过旧出问题

	global.Db = db //传参给项目全局变量global中的Db
}
