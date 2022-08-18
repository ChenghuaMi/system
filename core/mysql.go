package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"system/global"
	"fmt"
)

const DSN = "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
func LoadMysql() {
	var err error
	dsn := fmt.Sprintf(DSN,global.Config.Mysql.Username,global.Config.Mysql.Password,global.Config.Mysql.Host,global.Config.Mysql.Dbname)
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.Db = db
}