package models

import (
	"github.com/pibigstar/go-todo/models/db"
	"github.com/pibigstar/go-todo/utils/logger"
	"sync"
	"time"
)

var log = logger.New("models")

func init() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("init", err)
		}
	}()
	go func() {

		time.Sleep(time.Second * 10)
		once.Do(onceBody)
	}()
}

var once sync.Once
var onceBody = func() {
	if !db.OkMysqlInit {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Error("数据库初始化错误", err)
		}
	}()
	db.Mysql.Migrator().CreateTable(&User{})
	db.Mysql.Migrator().CreateIndex(&User{}, "id")
	db.Mysql.Migrator().CreateIndex(&User{}, "phone")
	db.Mysql.Migrator().CreateIndex(&User{}, "openId")

	db.Mysql.Migrator().CreateTable(&Admin{})
	db.Mysql.Migrator().CreateIndex(&Admin{}, "id")
	db.Mysql.Migrator().CreateIndex(&Admin{}, "username")

	db.Mysql.Migrator().CreateTable(&Group{})
	db.Mysql.Migrator().CreateIndex(&Group{}, "id")

	db.Mysql.Migrator().CreateTable(&Task{})
	db.Mysql.Migrator().CreateIndex(&Task{}, "id")
}
