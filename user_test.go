package workchatapp

import (
	"testing"
)

var testUser = User{
	Userid:     "laoji",
	Name:       "老纪",
	Department: []int32{1},
	Mobile:     "13121070248",
	Enable:     1,
}

func TestWorkChat_UserCreate(t *testing.T) {
	resp := testWorkChat.UserCreate(testUser)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
}
func TestWorkChat_UserGet(t *testing.T) {
	resp := testWorkChat.UserGet(testUser.Userid)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}
func TestWorkChat_UserUpdate(t *testing.T) {
	testUser.Enable = 0
	resp := testWorkChat.UserUpdate(testUser)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_UserSimpleList(t *testing.T) {
	resp := testWorkChat.UserSimpleList(1, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	for _, user := range resp.UserList {
		t.Log(user.UserId, user.Name)
	}
}
func TestWorkChat_UserList(t *testing.T) {
	resp := testWorkChat.UserList(1, 1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	for _, user := range resp.UserList {
		t.Log(user.Userid, user.Name)
	}
}
func TestWorkChat_UserId2OpenId(t *testing.T) {
	resp := testWorkChat.UserId2OpenId("jifengwei")
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_UserDelete(t *testing.T) {
	e := testWorkChat.UserDelete(testUser.Userid)
	if e.ErrCode != 0 {
		t.Error(e.ErrorMsg)
		return
	}
	t.Log(e)
}
func TestWorkChat_UserBatchDelete(t *testing.T) {
	e := testWorkChat.UserBatchDelete([]string{testUser.Userid})
	if e.ErrCode != 40031 {
		t.Error(e.ErrorMsg)
	}
	t.Log(e)
}
