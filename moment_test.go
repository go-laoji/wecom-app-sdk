package workchatapp

import (
	"testing"
)

var testMomentTask = MomentTask{
	Text: Text{
		Content: "测试发圈",
	},
}

var testJobId = ""

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
