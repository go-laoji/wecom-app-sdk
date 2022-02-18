package wecom

import (
	"fmt"
	"testing"
)

func TestWorkChat_AgentGet(t *testing.T) {
	resp := testWorkChat.AgentGet()
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_AgentList(t *testing.T) {
	resp := testWorkChat.AgentList()
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	if fmt.Sprintf("%v", resp.AgentList[0].AgentId) != testConfig.AppId {
		t.Error("Fail")
		return
	}
	t.Log(resp)
}
