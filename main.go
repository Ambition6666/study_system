package main

import (
	"fmt"
	"studysystem/clients"
	"studysystem/config"
	"studysystem/internal/http/router"
	"studysystem/internal/service/login"
	"studysystem/internal/service/pool"
	websokcet "studysystem/internal/service/websocket"
	"studysystem/logs"
	"studysystem/sql"
)

func init() {
	logs.InitLogger()
	config.ConfigInit()
	sql.InitSql()
	pool.P = pool.NewPool()
	login.W = *login.NewWorker(1, 1)
	websokcet.Manager = *websokcet.NewClientManager()
}
func main() {
	defer Close()
	jc, e := clients.InitJudgeGRPC()
	if e != nil {
		fmt.Println("初始化失败")
	}
	defer jc.Close()
	pc, e := clients.InitProGRPC()
	if e != nil {
		fmt.Println("初始化失败")
	}
	defer pc.Close()
	sql.RForm()
	go pool.P.Run()
	go websokcet.Manager.Start()
	r := router.InitRouter()
	r.Run(":" + config.ServerPort)
}
func Close() {

}
