package main

import (
	"work4/conf"
	"work4/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouters()
	_=r.Run(conf.HttpPort)
}