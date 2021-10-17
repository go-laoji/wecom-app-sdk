package workchatapp

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
	"os"
)

type Media struct {
	Type           string `json:"type" validate:"required,oneof=image voice video file"`
	AttachmentType int    `json:"attachment_type" validate:"required,oneof=1 2"`
	FilePath       string `json:"file_path" validate:"required"`
}

type MediaUploadResponse struct {
	internal.BizResponse
	Type     string `json:"type"`
	MediaId  string `json:"media_id"`
	CreateAt uint64 `json:"create_at"`
}

func (app workChat) MediaUploadAttachment(attrs Media) (resp MediaUploadResponse) {
	if ok := validate.Struct(attrs); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	if !isExists(attrs.FilePath) {
		resp.ErrCode = 500
		resp.ErrorMsg = fmt.Sprintf("%s 文件不存在！", attrs.FilePath)
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("media_type", attrs.Type)
	queryParams.Add("attachment_type", fmt.Sprintf("%v", attrs.AttachmentType))
	body, err := internal.HttpUploadMedia(
		fmt.Sprintf("/cgi-bin/media/upload_attachment?%s",
			queryParams.Encode()), attrs.FilePath)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
