package main

import (
	"studysystem/config"
	"studysystem/internal/http/router"
	"studysystem/internal/service/login"
	"studysystem/internal/service/pool"
	"studysystem/sql"
)

func init() {
	config.ConfigInit()
	sql.InitSql()
	pool.P = pool.NewPool()
	login.W = *login.NewWorker(1, 1)
}
func main() {
	sql.RForm()
	go pool.P.Run()
	r := router.InitRouter()
	r.Run(":" + config.ServerPort)
}
