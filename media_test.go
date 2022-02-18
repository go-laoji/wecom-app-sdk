package wecom

import "testing"

var testAttachment = Media{
	Type:           "image",
	AttachmentType: 1,
	FilePath:       "logo.png",
}

func TestWorkChat_MediaUpload(t *testing.T) {
	resp := testWorkChat.MediaUploadAttachment(testAttachment)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp)
}
