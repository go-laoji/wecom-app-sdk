package workchatapp

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

type ExternalMsg struct {
	ChatType       string                `json:"chat_type,omitempty" validate:"omitempty,oneof=single group"`
	ExternalUserid []string              `json:"external_userid,omitempty" validate:"required_without=Sender"`
	Sender         string                `json:"sender,omitempty" validate:"required_without=ExternalUserid"`
	Text           ExternalText          `json:"text,omitempty" validate:"required_without=Attachments"`
	Attachments    []ExternalAttachments `json:"attachments,omitempty" validate:"required_without=Text"`
}
type ExternalText struct {
	Content string `json:"content"`
}
type ExternalImage struct {
	MediaID string `json:"media_id" validate:"required_without=PicURL"`
	PicURL  string `json:"pic_url" validate:"required_without=MediaID"`
}
type ExternalLink struct {
	Title  string `json:"title" validate:"required"`
	Picurl string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	URL    string `json:"url"  validate:"required"`
}
type ExternalMiniprogram struct {
	Title      string `json:"title"  validate:"required"`
	PicMediaID string `json:"pic_media_id"  validate:"required"`
	Appid      string `json:"appid"  validate:"required"`
	Page       string `json:"page"  validate:"required"`
}
type ExternalVideo struct {
	MediaID string `json:"media_id" validate:"required"`
}
type ExternalFile struct {
	MediaID string `json:"media_id" validate:"required"`
}
type ExternalAttachments struct {
	Msgtype     string              `json:"msgtype" validate:"required"`
	Image       ExternalImage       `json:"image,omitempty"`
	Link        ExternalLink        `json:"link,omitempty"`
	Miniprogram ExternalMiniprogram `json:"miniprogram,omitempty"`
	Video       ExternalVideo       `json:"video,omitempty"`
	File        ExternalFile        `json:"file,omitempty"`
}

type AddMsgTemplateResponse struct {
	internal.BizResponse
	FailList []string `json:"fail_list"`
	MsgId    string   `json:"msgid"`
}

// AddMsgTemplate 创建企业群发
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92135
func (app workChat) AddMsgTemplate(msg ExternalMsg) (resp AddMsgTemplateResponse) {
	if ok := validate.Struct(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/add_msg_template?%s", queryParams.Encode()), msg)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GroupMsgListFilter struct {
	ChatType   string `json:"chat_type" validate:"required,oneof=single group"`
	StartTime  int64  `json:"start_time" validate:"required"`
	EndTime    int64  `json:"end_time" validate:"required"`
	Creator    string `json:"creator,omitempty"`
	FilterType int    `json:"filter_type,omitempty" validate:"omitempty,oneof=0 1 2"`
	Limit      int    `json:"limit" validate:"max=100"`
	Cursor     string `json:"cursor"`
}
type GetGroupMsgListV2Response struct {
	internal.BizResponse
	NextCursor   string         `json:"next_cursor"`
	GroupMsgList []GroupMsgList `json:"group_msg_list"`
}
type GroupMsgList struct {
	Msgid       string                `json:"msgid"`
	Creator     string                `json:"creator"`
	CreateTime  string                `json:"create_time"`
	CreateType  int                   `json:"create_type"`
	Text        ExternalText          `json:"text"`
	Attachments []ExternalAttachments `json:"attachments"`
}

// GetGroupMsgListV2 获取群发记录列表
//　参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/93338#%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%8F%91%E8%AE%B0%E5%BD%95%E5%88%97%E8%A1%A8
func (app workChat) GetGroupMsgListV2(filter GroupMsgListFilter) (resp GetGroupMsgListV2Response) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_groupmsg_list_v2?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GroupMsgTaskFilter struct {
	MsgId  string `json:"msgid" validate:"required"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type GetGroupMsgTaskResponse struct {
	internal.BizResponse
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		UserId   string `json:"userid"`
		Status   int    `json:"status"`
		SendTime int64  `json:"send_time"`
	} `json:"task_list"`
}

// GetGroupMsgTask 获取群发成员发送任务列表
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/93338#%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%8F%91%E6%88%90%E5%91%98%E5%8F%91%E9%80%81%E4%BB%BB%E5%8A%A1%E5%88%97%E8%A1%A8
func (app workChat) GetGroupMsgTask(filter GroupMsgTaskFilter) (resp GetGroupMsgTaskResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_groupmsg_task?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GroupMsgSendResultFilter struct {
	MsgId  string `json:"msgid" validate:"required"`
	UserId string `json:"userid" validate:"required"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type GetGroupMsgSendResultResponse struct {
	internal.BizResponse
	SendList []struct {
		ExternalUserId string `json:"external_user_id"`
		ChatId         string `json:"chat_id"`
		UserId         string `json:"userid"`
		Status         int    `json:"status"`
		SendTime       int64  `json:"send_time"`
	} `json:"send_list"`
}

// GetGroupMsgSendResult 获取企业群发成员执行结果
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/93338#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E7%BE%A4%E5%8F%91%E6%88%90%E5%91%98%E6%89%A7%E8%A1%8C%E7%BB%93%E6%9E%9C
func (app workChat) GetGroupMsgSendResult(filter GroupMsgSendResultFilter) (resp GetGroupMsgSendResultResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_groupmsg_send_result?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
