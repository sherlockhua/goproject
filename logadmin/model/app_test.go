package model

import (
	"testing"
)

func TestGetAllAppInfo(t *testing.T) {

	r, err := GetAllAppInfo()
	if err != nil {
		t.Fatalf("get all app info failed, err:%v", err)
		return
	}

	t.Logf("result:%v", r)
	return
}
