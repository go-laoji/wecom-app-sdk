package workchatapp

import (
	"testing"
	"time"
)

var testMomentTask = MomentTask{
	Text: Text{
		Content: "测试发圈",
	},
}

var testJobId = ""
var testMomentId = ""

func TestWorkChat_AddMomentTask(t *testing.T) {
	resp := testWorkChat.MediaUploadAttachment(testAttachment)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
	testMomentTask.Attachments = append(testMomentTask.Attachments, Attachments{
		Msgtype: "image",
		Image:   &Image{MediaID: resp.MediaId},
	})
	testMomentTask.VisibleRange = VisibleRange{SenderList: SenderList{UserList: []string{"jifengwei"}}}
	resp1 := testWorkChat.AddMomentTask(testMomentTask)
	if resp1.ErrCode != 0 {
		t.Error(resp1.ErrorMsg)
		return
	}
	t.Log(resp1)
	testJobId = resp1.JobId
}

func TestWorkChat_GetMomentTaskResult(t *testing.T) {
	resp := testWorkChat.GetMomentTaskResult(testJobId)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}

func TestWorkChat_GetMomentList(t *testing.T) {
	var testFilter = MomentListFilter{
		StartTime: time.Now().Unix() - 1000,
		EndTime:   time.Now().Unix()}
	resp := testWorkChat.GetMomentList(testFilter)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	for _, moment := range resp.MomentList {
		testMomentId = moment.MomentID
	}
	t.Log(resp)
}

func TestWorkChat_GetMomentTask(t *testing.T) {
	var testFilter = MomentTaskFilter{
		MomentId: testMomentId,
	}
	resp := testWorkChat.GetMomentTask(testFilter)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}
