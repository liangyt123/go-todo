package test

import (
	"testing"

	"github.com/liangyt123/go-todo/utils"
)

func TestMd5(t *testing.T) {
	t.Log(utils.Md5("liangyt123"))
}
