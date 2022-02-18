package wecom

import (
	"testing"
	"time"
)

var testContactMe = ContactMe{
	Type:  1,
	Scene: 2,
	User:  []string{"jifengwei"},
	State: "单元测试联系我",
}

func TestWorkChat_ExternalAddContactWay(t *testing.T) {
	resp := testWorkChat.ExternalAddContactWay(testContactMe)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	testContactMe.ConfigId = resp.ConfigId
	t.Log(resp)
}

func TestWorkChat_ExternalUpdateContactWay(t *testing.T) {
	testContactMe.Remark = "更新备注"
	resp := testWorkChat.ExternalUpdateContactWay(testContactMe)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_ExternalGetContactWay(t *testing.T) {
	resp := testWorkChat.ExternalGetContactWay(testContactMe.ConfigId)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_ExternalListContactWay(t *testing.T) {
	resp := testWorkChat.ExternalListContactWay(1633888980, time.Now().Unix(), "", 100)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	for _, me := range resp.ContactWay {
		t.Log(me.ConfigId)
	}
}

func TestWorkChat_ExternalDeleteContactWay(t *testing.T) {
	resp := testWorkChat.ExternalDeleteContactWay(testContactMe.ConfigId)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_ExternalCloseTempChat(t *testing.T) {
	//	TODO:
}
