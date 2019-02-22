package gotest

import (
	"testing"
)

func TestDivision(t *testing.T) {
	_, err := Division(10, 0)
	if err == nil {
		t.Error("0除异常没有处理！")
	}
	res, err := Division(10, 2)
	if err != nil {
		t.Error("除法异常！", err)
	}

	if res != 5 {
		t.Error("除法处理逻辑错误！")
	}
}
