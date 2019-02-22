package gotest

import (
	"testing"
)

func TestDivision1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没通过")
	} else {
		t.Log("第一个测试通过了！")
	}
}

func TestDivision2(t *testing.T) {
	t.Error("失败的用例！")
	// res, err := Division(10, 0)
	// if err != nil {
	// 	t.Error(err)
	// }
	// t.Log("result: ", res)
}
