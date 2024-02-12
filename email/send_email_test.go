package email

import (
	"testing"
)

func TestSendUtil_IsCCOrBCCSend(t *testing.T) {
	err := NewSimpleSendUtil("228664584@qq.com",
		"228664584@qq.com",
		"test",
		"test",
		nil,
		&User{
			Identity: "",
			Username: "228664584@qq.com",
			Password: "usuzpuerbdddbgeb",
		},
		QQ).IsCCOrBCCSend(true, false)

	if err != nil {
		t.Error(err)
	}
}

func TestSendUtil_SimpleSend(t *testing.T) {
	util := NewCCSendUtil("228664584@qq.com",
		"228664584@qq.com",
		"test",
		"test",
		nil,
		&User{
			Identity: "",
			Username: "228664584@qq.com",
			Password: "usuzpuerbdddbgeb",
		},
		[]string{"228664584@qq.com", "427945607@qq.com"},
		QQ)
	err := util.IsCCOrBCCSend(true, true)

	if err != nil {
		t.Error(err)
	}

}
