package main

import (
	"fraise/config"
	"fraise/net"
	"fraise/server"
)

func main() {
	config.InitConfig()
	server.InitList()
	net.InitHttp()
}
