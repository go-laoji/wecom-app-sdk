package workchatapp

//　参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92572

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

type ConclusionsText struct {
	Content string `json:"content"`
}

type ConclusionsImage struct {
	MediaId string `json:"media_id"`
	PicUrl  string `json:"pic_url"`
}

type ConclusionsLink struct {
	Title  string `json:"title"`
	PicUrl string `json:"pic_url"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
}

type ConclusionsMiniProgram struct {
	Title      string `json:"title"`
	PicMediaId string `json:"pic_media_id"`
	AppId      string `json:"appid"`
	Page       string `json:"page"`
}

type ContactMe struct {
	ConfigId      string   `json:"config_id,omitempty"`
	Type          int      `json:"type" validate:"required,oneof=1 2"`
	Scene         int      `json:"scene" validate:"required,oneof=1 2"`
	Style         int      `json:"style"`
	Remark        string   `json:"remark"`
	SkipVerify    bool     `json:"skip_verify"`
	State         string   `json:"state"`
	User          []string `json:"user"`
	Party         []int32  `json:"party"`
	IsTemp        bool     `json:"is_temp"`
	ExpiresIn     int32    `json:"expires_in"`
	ChatExpiresIn int32    `json:"chat_expires_in"`
	UnionId       string   `json:"unionid"`
	Conclusions   struct {
		*ConclusionsText        `json:"text,omitempty"`
		*ConclusionsImage       `json:"image,omitempty"`
		*ConclusionsLink        `json:"link,omitempty"`
		*ConclusionsMiniProgram `json:"miniprogram,omitempty"`
	} `json:"conclusions"`
}

type ContactMeAddResponse struct {
	internal.BizResponse
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

// ExternalAddContactWay 配置客户联系「联系我」方式
func (app workChat) ExternalAddContactWay(me ContactMe) (resp ContactMeAddResponse) {
	if ok := validate.Struct(me); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/add_contact_way?%s", queryParams.Encode()), me)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// ExternalUpdateContactWay 更新企业已配置的「联系我」方式
func (app workChat) ExternalUpdateContactWay(me ContactMe) (resp internal.BizResponse) {
	if ok := validate.Struct(me); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/update_contact_way?%s", queryParams.Encode()), me)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ContactMeGetResponse struct {
	internal.BizResponse
	ContactWay struct {
		ConfigId string `json:"config_id"`
		ContactMe
	} `json:"contact_way"`
}

// ExternalGetContactWay 获取企业已配置的「联系我」方式
func (app workChat) ExternalGetContactWay(configId string) (resp ContactMeGetResponse) {
	p := H{"config_id": configId}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_contact_way?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ContactMeListResponse struct {
	internal.BizResponse
	ContactWay []struct {
		ConfigId string `json:"config_id"`
	} `json:"contact_way"`
	NextCursor string `json:"next_cursor"`
}

// ExternalListContactWay 获取企业配置的「联系我」二维码和「联系我」小程序插件列表。不包含临时会话。
// 注意，该接口仅可获取2021年7月10日以后创建的「联系我」
func (app workChat) ExternalListContactWay(startTime, endTime int64, cursor string, limit int) (resp ContactMeListResponse) {
	p := H{"start_time": startTime, "end_time": endTime, "cursor": cursor, "limit": limit}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/list_contact_way?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// ExternalDeleteContactWay 删除企业已配置的「联系我」方式
func (app workChat) ExternalDeleteContactWay(configId string) (resp internal.BizResponse) {
	p := H{"config_id": configId}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/del_contact_way?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// ExternalCloseTempChat 结束临时会话
func (app workChat) ExternalCloseTempChat(userId, externalUserId string) (resp internal.BizResponse) {
	p := H{"userid": userId, "external_userid": externalUserId}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/del_contact_way?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
