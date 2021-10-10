package workchatapp

import "testing"

var testTag = Tag{
	TagId:   10,
	TagName: "test tag",
}

func TestWorkChat_TagCreate(t *testing.T) {
	resp := testWorkChat.TagCreate(testTag)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_TagUpdate(t *testing.T) {
	testTag.TagName = "测试标签"
	resp := testWorkChat.TagUpdate(testTag)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_TagUserList(t *testing.T) {
	resp := testWorkChat.TagUserList(testTag.TagId)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_TagAddUsers(t *testing.T) {
	resp := testWorkChat.TagAddUsers(testTag.TagId, []string{"jifengwei"}, []int32{1})
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_TagDelUsers(t *testing.T) {
	resp := testWorkChat.TagDelUsers(testTag.TagId, []string{"jifengwei"}, []int32{1})
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_TagDelete(t *testing.T) {
	resp := testWorkChat.TagDelete(testTag.TagId)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_TagList(t *testing.T) {
	resp := testWorkChat.TagList()
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}
