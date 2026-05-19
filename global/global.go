package global

import "gorm.io/gorm"

var (
	Db *gorm.DB //设置整个项目的全局变量
)
