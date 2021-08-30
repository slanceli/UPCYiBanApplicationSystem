package main

import (
	"UPCYiBanApplicationSystemGo/dao"
	"UPCYiBanApplicationSystemGo/router"
)

func main() {
	dao.DBInit(true)
	router.InitRouter()
	router.RunRouter()
}