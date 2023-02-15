package models

import (
	"github.com/liangyt123/go-todo/models/db"
	"time"
)

var MTaskUser = &TaskUser{}

type TaskUser struct {
	ID         int
	TaskID     int       `gorm:"column:task_id"`
	UserID     string    `gorm:"column:task_id"`
	Status     int       `gorm:"column:status"`
	CreateTime time.Time `gorm:"column:create_time"`
	IsDelete   bool      `gorm:"column:is_delete"`
}

func (*TaskUser) TableName() string {
	return "task_user"
}

func (t *TaskUser) Create(taskUser *TaskUser) error {
	return db.Mysql.Table(t.TableName()).Create(taskUser).Error
}
