package workchatapp

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

type MomentTask struct {
	Text         Text          `json:"text,omitempty"`
	Attachments  []Attachments `json:"attachments" validate:"required"`
	VisibleRange VisibleRange  `json:"visible_range,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}
type Image struct {
	MediaID string `json:"media_id" validate:"required"`
}
type Video struct {
	MediaID string `json:"media_id" validate:"required"`
}
type Link struct {
	Title   string `json:"title"`
	URL     string `json:"url" validate:"required"`
	MediaID string `json:"media_id" validate:"required"`
}
type Attachments struct {
	Msgtype string `json:"msgtype" validate:"required,oneof=image link video"`
	Image   *Image `json:"image,omitempty" validate:"required_without_all=Video Link"`
	Video   *Video `json:"video,omitempty" validate:"required_without_all=Image Link"`
	Link    *Link  `json:"link,omitempty" validate:"required_without_all=Video Image"`
}
type SenderList struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int    `json:"department_list"`
}
type ExternalContactList struct {
	TagList []string `json:"tag_list"`
}
type VisibleRange struct {
	SenderList          SenderList          `json:"sender_list,omitempty"`
	ExternalContactList ExternalContactList `json:"external_contact_list,omitempty"`
}

type AddMomentTaskResponse struct {
	internal.BizResponse
	JobId string `json:"jobid"`
}

// AddMomentTask 创建发表任务
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/95094#%E5%88%9B%E5%BB%BA%E5%8F%91%E8%A1%A8%E4%BB%BB%E5%8A%A1
func (app workChat) AddMomentTask(task MomentTask) (resp AddMomentTaskResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/add_moment_task?%s", queryParams.Encode()), task)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetMomentTaskResultResponse struct {
	internal.BizResponse
	Status int    `json:"status"`
	Type   string `json:"type"`
	Result struct {
		internal.BizResponse
		MomentId          string `json:"moment_id"`
		InvalidSenderList struct {
			UserList       []string `json:"user_list"`
			DepartmentList []int32  `json:"department_list"`
		} `json:"invalid_sender_list"`
		InvalidExternalContactList struct {
			TagList []string `json:"tag_list"`
		} `json:"invalid_external_contact_list"`
	}
}

// GetMomentTaskResult 获取任务创建结果
func (app workChat) GetMomentTaskResult(jobId string) (resp GetMomentTaskResultResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("jobid", jobId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_task_result?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
