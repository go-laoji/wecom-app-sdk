package wecom

import "testing"

func TestWorkChat_GroupChatList(t *testing.T) {
	var gcf = GroupChatListFilter{
		Limit: 10,
	}
	resp := testWorkChat.GroupChatList(gcf)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_GroupChat(t *testing.T) {

}

func TestWorkChat_GroupOpenId2ChatId(t *testing.T) {

}
