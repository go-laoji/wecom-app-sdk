package workchatapp

import (
	"testing"
)

func TestWorkChat_ExternalContactGetFollowUserList(t *testing.T) {
	resp := testWorkChat.ExternalContactGetFollowUserList()
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}
