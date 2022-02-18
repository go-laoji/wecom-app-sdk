package wecom

import "testing"

var testDepart = Department{
	Id:       5,
	Order:    1,
	ParentId: 1,
	Name:     "技术委员会",
}

func TestWorkChat_DepartmentCreate(t *testing.T) {
	resp := testWorkChat.DepartmentCreate(testDepart)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_DepartmentUpdate(t *testing.T) {
	testDepart.NameEn = "Technical Committee"
	resp := testWorkChat.DepartmentUpdate(testDepart)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_DepartmentList(t *testing.T) {
	resp := testWorkChat.DepartmentList(1)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_DepartmentDelete(t *testing.T) {
	resp := testWorkChat.DepartmentDelete(testDepart.Id)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}
