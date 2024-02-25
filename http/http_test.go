/*
@author: ledger
@since: 2024/2/23
*/
package test

import (
	"testing"
)

type param struct {
	Name string
	Age  int
}

func TestGet(t *testing.T) {
	get, err := NewHttpDos("https://www.baidu.com", nil, nil, nil).Get()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(get))
}

func TestGet2(t *testing.T) {

}
