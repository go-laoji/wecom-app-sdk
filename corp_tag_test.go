package wecom

import (
	"testing"
)

var testCorpTagGroup = CorpTagGroup{
	GroupName: "测试标签组",
	Order:     1,
	Tag: []CorpTag{
		{Name: "标签一", Order: 1},
		{Name: "标签二", Order: 2},
		{Name: "标签三", Order: 3},
	},
}

var testUpdateTagId = ""

func TestWorkChat_CorpTagList(t *testing.T) {
	resp := testWorkChat.CorpTagList([]string{}, []string{})
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
}

func TestWorkChat_CorpTagAdd(t *testing.T) {
	resp := testWorkChat.CorpTagAdd(testCorpTagGroup)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	for _, tag := range resp.TagGroup.Tag {
		testUpdateTagId = tag.Id
		break
	}
	testCorpTagGroup = resp.TagGroup
	t.Log(resp)
}

func TestWorkChat_CorpTagUpdate(t *testing.T) {
	testCorpTag := CorpTag{Id: testUpdateTagId, Name: "测试标签1", Order: 3}
	resp := testWorkChat.CorpTagUpdate(testCorpTag)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_CorpTagDelete(t *testing.T) {
	resp := testWorkChat.CorpTagDelete([]string{testUpdateTagId}, []string{})
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_CorpTagList2(t *testing.T) {
	resp := testWorkChat.CorpTagList([]string{}, []string{testCorpTagGroup.GroupId})
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	for _, tag := range resp.TagGroup[0].Tag {
		if tag.Id != testUpdateTagId {
			continue
		}
		if tag.Deleted {
			t.Log("标签删除成功")
		} else {
			t.Error("删除标签失败")
		}
	}
}

func TestWorkChat_CorpTagDelete2(t *testing.T) {
	resp := testWorkChat.CorpTagDelete([]string{}, []string{testCorpTagGroup.GroupId})
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}
