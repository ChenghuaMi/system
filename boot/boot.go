package boot

import (
	"system/app"
	"system/core"
)

func init() {
	core.Viper()
	app.InitRoute()
	core.LoadMysql()
}

