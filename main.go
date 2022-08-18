package main

import (
	_ "system/boot"
	"system/route"
)

func main() {
	route.InitRoute().Run()
}
