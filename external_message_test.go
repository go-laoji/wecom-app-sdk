package wecom

import (
	"testing"
	"time"
)

var testMsg = ExternalMsg{
	Text: ExternalText{
		Content: "test",
	},
	Sender: "jifengwei",
}
var testMsgId = ""

func TestWorkChat_AddMsgTemplate(t *testing.T) {
	resp := testWorkChat.AddMsgTemplate(testMsg)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_GetGroupMsgListV2(t *testing.T) {
	var filter = GroupMsgListFilter{
		ChatType:  "single",
		StartTime: time.Now().Unix() - 1000,
		EndTime:   time.Now().Unix(),
	}
	resp := testWorkChat.GetGroupMsgListV2(filter)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	for _, msg := range resp.GroupMsgList {
		testMsgId = msg.Msgid
	}
	t.Log(resp)
}

func TestWorkChat_GetGroupMsgTask(t *testing.T) {
	var filter = GroupMsgTaskFilter{
		MsgId: testMsgId,
	}
	resp := testWorkChat.GetGroupMsgTask(filter)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_GetGroupMsgSendResult(t *testing.T) {
	var filter = GroupMsgSendResultFilter{
		MsgId:  testMsgId,
		UserId: "jifengwei",
	}
	resp := testWorkChat.GetGroupMsgSendResult(filter)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}
