package test

import (
	"testing"

	"github.com/liangyt123/go-todo/utils"
)

func TestUUID(t *testing.T) {
	uid := utils.GetUUID()
	t.Log(uid)
}
