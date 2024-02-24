package email

import (
	"testing"
)

func TestSendUtil_IsCCOrBCCSend(t *testing.T) {
	err := NewSimpleSendUtil("xxx@qq.com",
		"xxx@qq.com",
		"test",
		"test",
		nil,
		&User{
			Identity: "",
			Username: "xxx@qq.com",
			Password: "xxx",
		},
		QQ).IsCCOrBCCSend(true, false)

	if err != nil {
		t.Error(err)
	}
}

func TestSendUtil_SimpleSend(t *testing.T) {
	util := NewCCSendUtil("xxx@qq.com",
		"xxx@qq.com",
		"test",
		"test",
		nil,
		&User{
			Identity: "",
			Username: "xxx@qq.com",
			Password: "xxx",
		},
		[]string{"xxx@qq.com", "xxx@qq.com"},
		QQ)
	err := util.IsCCOrBCCSend(true, true)

	if err != nil {
		t.Error(err)
	}

}
