package test

import (
	"github.com/liangyt123/go-todo/utils"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	t.Log(utils.TimeFormat(time.Now()))
}
