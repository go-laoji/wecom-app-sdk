package wecom

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-app-sdk/internal"
)

type MomentTask struct {
	Text         Text          `json:"text,omitempty"`
	Attachments  []Attachments `json:"attachments" validate:"required_without=Text.Content"`
	VisibleRange VisibleRange  `json:"visible_range,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}
type Image struct {
	MediaID string `json:"media_id" validate:"required"`
}

// Video 应用消息发关时title和description为可选项
// 朋友圈发送时只设置　media_id即可
type Video struct {
	MediaID     string `json:"media_id" validate:"required"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
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
func (app weCom) AddMomentTask(task MomentTask) (resp AddMomentTaskResponse) {
	if ok := validate.Struct(task); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
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
func (app weCom) GetMomentTaskResult(jobId string) (resp GetMomentTaskResultResponse) {
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

type MomentListFilter struct {
	StartTime  int64  `json:"start_time" validate:"required"`
	EndTime    int64  `json:"end_time" validate:"required"`
	Creator    string `json:"creator,omitempty"`
	FilterType int    `json:"filter_type,omitempty" validate:"omitempty,oneof=0 1 2"`
	Cursor     string `json:"cursor"`
	Limit      int    `json:"limit"`
}

type GetMomentListResponse struct {
	internal.BizResponse
	NextCursor string       `json:"next_cursor"`
	MomentList []MomentList `json:"moment_list"`
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
}
type MomentList struct {
	MomentID    string   `json:"moment_id"`
	Creator     string   `json:"creator"`
	CreateTime  string   `json:"create_time"`
	CreateType  int      `json:"create_type"`
	VisibleType int      `json:"visible_type"`
	Text        Text     `json:"text"`
	Image       []Image  `json:"image"`
	Video       Video    `json:"video"`
	Link        Link     `json:"link"`
	Location    Location `json:"location"`
}

func (app weCom) GetMomentList(filter MomentListFilter) (resp GetMomentListResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_list?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type MomentTaskFilter struct {
	MomentId string `json:"moment_id" validate:"required"`
	Cursor   string `json:"cursor"`
	Limit    int    `json:"limit"`
}

type GetMomentTaskResponse struct {
	internal.BizResponse
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		UserId        string `json:"userid"`
		PublishStatus int    `json:"publish_status"`
	} `json:"task_list"`
}

func (app weCom) GetMomentTask(filter MomentTaskFilter) (resp GetMomentTaskResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_task?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
