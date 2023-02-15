package test

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/test/gtest"
	"testing"

	"github.com/liangyt123/go-todo/models"
)

func TestCreateUser(t *testing.T) {
	user := &models.User{
		OpenID:   "liangyt123",
		NickName: "派大星",
	}
	err := user.Create(user)
	gtest.Assert(err, nil)

	getUser, err := models.MUser.GetUserByOpenID("liangyt123")
	gtest.Assert(err, nil)
	glog.Print(getUser)

}
