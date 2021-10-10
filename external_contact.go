package workchatapp

// 客户管理

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

type ExternalContactGetFollowUserListResponse struct {
	internal.BizResponse
	FollowUser []string `json:"follow_user"`
}

// ExternalContactGetFollowUserList 获取配置了客户联系功能的成员列表
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92571
func (app workChat) ExternalContactGetFollowUserList() (resp ExternalContactGetFollowUserListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/externalcontact/get_follow_user_list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ExternalContactListResponse struct {
	internal.BizResponse
	ExternalUserId []string `json:"external_userid"`
}

// ExternalContactList 获取客户列表
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92113
func (app workChat) ExternalContactList(userId string) (resp ExternalContactListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("userid", userId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/externalcontact/list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ExternalContact struct {
	ExternalUserId  string `json:"external_userid"`
	Name            string `json:"name"`
	Position        string `json:"position"`
	Avatar          string `json:"avatar"`
	CorpName        string `json:"corp_name"`
	CorpFullName    string `json:"corp_full_name"`
	Type            int    `json:"type"`
	Gender          int    `json:"gender"`
	UnionId         string `json:"unionid"`
	ExternalProfile struct {
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			MiniProgram struct {
				AppId    string `json:"appid"`
				PagePath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		}
	} `json:"external_profile"`
}

type FollowUser struct {
	UserId      string `json:"userid"`
	Remark      string `json:"remark,omitempty"`
	Description string `json:"description,omitempty"`
	CreateTime  int64  `json:"createtime"`
	Tags        []struct {
		GroupName string `json:"group_name"`
		TagName   string `json:"tag_name"`
		TagId     string `json:"tag_id"`
		Type      int    `json:"type"`
	} `json:"tags,omitempty"`
	RemarkCorpName string   `json:"remark_corp_name,omitempty"`
	RemarkMobiles  []string `json:"remark_mobiles,omitempty"`
	State          string   `json:"state,omitempty"`
	OperUserId     string   `json:"oper_userid,omitempty"`
	AddWay         int      `json:"add_way,omitempty"`
}

type ExternalContactGetResponse struct {
	internal.BizResponse
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
	NextCursor      string          `json:"next_cursor"`
}

// ExternalContactGet 获取客户详情
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92114
// 当客户在企业内的跟进人超过500人时需要使用cursor参数进行分页获取
func (app workChat) ExternalContactGet(externalUserId, cursor string) (resp ExternalContactGetResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("external_userid", externalUserId)
	queryParams.Add("cursor", cursor)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/externalcontact/get?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ExternalContactBatchGetByUserResponse struct {
	internal.BizResponse
	ExternalContactList []struct {
		ExternalContact ExternalContact `json:"external_contact"`
		FollowInfo      FollowUser      `json:"follow_info"`
	} `json:"external_contact_list"`
	NextCursor string `json:"next_cursor"`
}

// ExternalContactBatchGetByUser 批量获取客户详情
// 企业可通过此接口获取指定成员添加的客户信息列表。
// 参考连接 https://open.work.weixin.qq.com/api/doc/90000/90135/92994
func (app workChat) ExternalContactBatchGetByUser(userIds []string, cursor string, limit int) (resp ExternalContactBatchGetByUserResponse) {
	p := H{"userid_list": userIds, "cursor": cursor, "limit": limit}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/batch/get_by_user?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ExternalContactRemarkRequest struct {
	UserId           string   `json:"user_id" validate:"required"`
	ExternalUserid   string   `json:"external_userid" validate:"required"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaId string   `json:"remark_pic_mediaid"`
}

// ExternalContactRemark 修改客户备注信息
// 参考连接 https://open.work.weixin.qq.com/api/doc/90000/90135/92115
func (app workChat) ExternalContactRemark(remark ExternalContactRemarkRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(remark); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/remark?%s", queryParams.Encode()), remark)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
