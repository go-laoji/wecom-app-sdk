package wecom

import (
	"testing"
)

//
// 涉及到外部联系人的id不好获取，单元测试不完整
//

func TestWorkChat_ExternalContactGetFollowUserList(t *testing.T) {
	resp := testWorkChat.ExternalContactGetFollowUserList()
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_ExternalContactList(t *testing.T) {

}

func TestWorkChat_ExternalContactGet(t *testing.T) {

}

func TestWorkChat_ExternalContactBatchGetByUser(t *testing.T) {

}

func TestWorkChat_ExternalContactRemark(t *testing.T) {

}
