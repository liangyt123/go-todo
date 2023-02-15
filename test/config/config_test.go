package test

import (
	"testing"

	"github.com/liangyt123/go-todo/config"
)

func TestConfig(t *testing.T) {
	t.Logf("%+v", config.ServerConfig)
}
