package model

import (
	"system/core"
	"system/global"
	"testing"
	"fmt"
)
func init() {
	core.Viper()
	core.LoadMysql()
}

func TestAddAdminRole(t *testing.T) {

	roles := []AdminRole{
		{
			RoleName: "演员",
			Description: "演员",
			System: 1,
			Disabled: 1,
		},
		{
			RoleName: "管理员",
			Description: "管理员",
			System: 0,
			Disabled: 1,
		},

	}
	for _,role := range roles {
		result := global.Db.Create(&role)
		fmt.Println(result.Error,result.RowsAffected)
	}
}