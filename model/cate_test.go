package model

import (
	"system/core"
	"system/global"
	"testing"
	"time"
)

func init() {
	core.Viper()
	core.LoadMysql()
}
func TestInsertCate(t *testing.T) {
	data := []Cate{
		{CateName: "服装",CreatedAt: time.Now(),Brand: Brand{
			BrandName: "李宁",
		}},
		{CateName: "电脑",Brand: Brand{
			BrandName: "苹果",
		}},
	}
	global.Db.Create(&data)
}
