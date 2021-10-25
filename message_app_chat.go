package workchatapp

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

//企业微信支持企业自建应用通过接口创建群聊并发送消息到群，让重要的消息可更及时推送给群成员，方便协同处理。
//应用消息仅限于发送到通过接口创建的内部群聊，不支持添加企业外部联系人进群。此接口暂时仅支持企业内接入使用。

type AppChatCreateRequest struct {
	Name     string   `json:"name,omitempty"`
	Owner    string   `json:"owner,omitempty"`
	UserList []string `json:"userlist" validate:"required,min=2,max=2000"`
	ChatId   string   `json:"chatid,omitempty"`
}

type AppChatCreateResponse struct {
	internal.BizResponse
	ChatId string `json:"chatid"`
}

// AppChatCreate 创建群聊会话
// https://open.work.weixin.qq.com/api/doc/90000/90135/90245
func (app workChat) AppChatCreate(request AppChatCreateRequest) (resp AppChatCreateResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/appchat/create?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type AppChatUpdateRequest struct {
	ChatId      string   `json:"chatid" validate:"required"`
	Name        string   `json:"name,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	AddUserList []string `json:"add_user_list,omitempty"`
	DelUserList []string `json:"del_user_list,omitempty"`
}

// AppChatUpdate 修改群聊会话
// https://open.work.weixin.qq.com/api/doc/90000/90135/90246
func (app workChat) AppChatUpdate(request AppChatUpdateRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/appchat/update?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type AppChatGetResponse struct {
	internal.BizResponse
	ChatInfo struct {
		ChatId   string   `json:"chatid"`
		Name     string   `json:"name"`
		Owner    string   `json:"owner"`
		UserList []string `json:"userlist"`
	} `json:"chat_info"`
}

// AppChatGet 获取群聊会话
// https://open.work.weixin.qq.com/api/doc/90000/90135/90247
func (app workChat) AppChatGet(chatId string) (resp AppChatGetResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("chatid", chatId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/appchat/update?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// AppChatSend 应用推送消息
// https://open.work.weixin.qq.com/api/doc/90000/90135/90248
func (app workChat) AppChatSend(msg interface{}, chatId string) (resp internal.BizResponse) {
	if ok := validate.Struct(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	h := H{}
	buf, _ := json.Marshal(msg)
	json.Unmarshal(buf, &h)
	h["chatid"] = chatId
	switch msg.(type) {
	case TextMessage:
		h["msgtype"] = "text"
	case ImageMessage:
		h["msgtype"] = "image"
	case VoiceMessage:
		h["msgtype"] = "voice"
	case VideoMessage:
		h["msgtype"] = "video"
	case FileMessage:
		h["msgtype"] = "file"
	case TextCardMessage:
		h["msgtype"] = "textcard"
	case NewsMessage:
		h["msgtype"] = "news"
	case MpNewsMessage:
		h["msgtype"] = "mpnews"
	case MarkDownMessage:
		h["msgtype"] = "markdown"
	default:
		resp.ErrCode = 500
		resp.ErrorMsg = "不支持的消息类型"
		return
	}

	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/appchat/send?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
