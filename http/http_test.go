/*
@author: ledger
@since: 2024/2/23
*/
package test

import (
	"testing"
)

func TestHttpDos_Get(t *testing.T) {
	get, err := NewHttpDos("https://www.baidu.com", nil, nil).Get()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(get))
}

type param struct {
	Name string
	Age  int
}

func TestGet(t *testing.T) {
	p := param{
		Name: "",
		Age:  0,
	}
	newUrl, err := MontageURL("https://www.baidu.com", p)
	if err != nil {
		t.Error(err)
	}
	get, err := NewHttpDos(newUrl, nil, nil).Get()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(get))
}

func TestGet2(t *testing.T) {

}
