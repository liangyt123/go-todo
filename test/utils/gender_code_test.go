package test

import (
	"github.com/liangyt123/go-todo/utils"
	"testing"
)

func TestGenderCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(utils.GenderCode())
	}
}
