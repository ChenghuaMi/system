package global

import (
	"system/config"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db *gorm.DB
)
