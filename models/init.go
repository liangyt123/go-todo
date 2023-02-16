package models

import (
	"github.com/liangyt123/go-todo/models/db"
	"github.com/pibigstar/go-todo/utils/logger"
)

var log = logger.New("models")

var InitTable = func() {
	if !db.OkMysqlInit {
		return
	}
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
