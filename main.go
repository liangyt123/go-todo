package main

import (
	_ "github.com/liangyt123/go-todo/controller"
	_ "github.com/liangyt123/go-todo/controller/admin"

	"github.com/liangyt123/go-todo/models"
	"github.com/liangyt123/go-todo/models/db"

	"github.com/gogf/gf/frame/g"
	"github.com/liangyt123/go-todo/config"
)

func main() {
	s := g.Server()
	port := config.ServerConfig.Port
	s.SetPort(port)
	host := config.ServerConfig.Host
	s.Domain(host)
	// 开启日志
	s.SetLogPath("log/todo.log")
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	// 开启https
	s.EnableHTTPS("https/ssl.pem", "https/ssl.key")
	s.SetHTTPSPort(443)
	// 开启性能分析，可访问页面/debug/pprof
	s.EnablePprof()
	db.Init1()
	models.InitTable()
	s.Run()
}
