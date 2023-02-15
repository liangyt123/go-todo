package models

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/go-gorm/gorm"
	"time"
	"github.com/pibigstar/go-todo/models/db"
)

// MGroup 引用
var MGroup = &Group{}

// Group 组织表
type Group struct {
	ID            int       `gorm:"column:id;primary_key"`
	GroupName     string    `gorm:"column:group_name"`
	GroupDescribe string    `gorm:"column:group_describe"`
	GroupMaster   string    `gorm:"column:group_master"`
	GroupCode     string    `gorm:"column:group_code"`
	CreateUser    string    `gorm:"column:create_user"`
	CreateTime    time.Time `gorm:"column:create_time"`
	UpdateTime    time.Time `gorm:"column:update_time"`
	IsDelete      bool      `gorm:"column:is_delete"`
	JoinMethod    string    `gorm:"column:join_method"`
	Question      string    `gorm:"column:question"`
	Answer        string    `gorm:"column:answer"`
}

// TableName 组织表
func (Group) TableName() string {
	return "groups"
}




// Insert 创建
func (*Group) Create(group *Group) error {
	return db.Mysql.Table("groups").Create(&group).Error
}

// GetGroupByID 根据ID获取组织
func (t *Group) GetGroupByID(groupID int) (*Group, error) {
	var group Group
	err := db.Mysql.Table(t.TableName()).Where("id = ?", groupID).First(&group).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Info("没有此组织", "groupId", gconv.String(groupID))
		}
		return nil, err
	}
	return &group, nil
}

// GetGroupsByUserID 获取用户创建的组织
func (t *Group) GetUserCreateGroups(openID string) (*[]Group, error) {
	var groups []Group
	err := db.Mysql.Table(t.TableName()).Where("create_user = ?", openID).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return &groups, nil
}
func (t *Group) ListGroup() (*[]Group, error) {
	var groups []Group
	err := db.Mysql.Table(t.TableName()).Where("is_delete = ?", false).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return &groups, nil
}
func (t *Group) GroupDelete(id int) error {
	group := Group{
		ID: id,
	}
	err := db.Mysql.Table(t.TableName()).Delete(&group).Error
	return err
}
