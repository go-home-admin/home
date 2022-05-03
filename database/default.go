package database

import "gorm.io/gorm"

// DB 一个全局的默认链接
var DB func() *gorm.DB
