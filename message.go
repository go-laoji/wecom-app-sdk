package workchatapp

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

type Message struct {
	ToUser                 string `json:"touser,omitempty" validate:"omitempty,required_without=ToParty ToTag"`
	ToParty                string `json:"toparty,omitempty" validate:"omitempty,required_without=ToUser ToTag"`
	ToTag                  string `json:"totag,omitempty" validate:"omitempty,required_without=ToParty ToUser"`
	EnableIDTrans          int    `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval,omitempty"`
}

type TextMessage struct {
	Message
	Safe int  `json:"safe,omitempty" validate:"omitempty,oneof=0 1"`
	Text Text `json:"text" validate:"required"`
}

type ImageMessage struct {
	Message
	Safe  int        `json:"safe,omitempty" validate:"omitempty,oneof=0 1"`
	Image MultiMedia `json:"image" validate:"required"`
}
type MultiMedia struct {
	MediaId string `json:"media_id" validate:"required"`
}
type VoiceMessage struct {
	Message
	Safe  int        `json:"safe,omitempty"`
	Voice MultiMedia `json:"voice" validate:"required"`
}

type VideoMessage struct {
	Message
	Safe  int   `json:"safe,omitempty" validate:"omitempty,oneof=0 1"`
	Video Video `json:"video" validate:"required"`
}

type FileMessage struct {
	Message
	Safe int        `json:"safe,omitempty"`
	File MultiMedia `json:"file" validate:"required"`
}

type TextCard struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Url         string `json:"url" validate:"required"`
	BtnTxt      string `json:"btntxt"`
}
type TextCardMessage struct {
	Message
	TextCard TextCard `json:"textcard" validate:"required"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
	AppId       string `json:"appid" validate:"required_without=Url,required_with=PagePath"`
	PagePath    string `json:"pagepath" validate:"required_with=AppId"`
}
type News struct {
	Articles []Article `json:"articles" validate:"required,max=8"`
}
type NewsMessage struct {
	Message
	News News `json:"news" validate:"required"`
}

type MpArticle struct {
	Title            string `json:"title" validate:"required"`
	ThumbMediaId     string `json:"thumb_media_id" validate:"required"`
	Author           string `json:"author,omitempty"`
	ContentSourceUrl string `json:"content_source_url,omitempty"`
	Content          string `json:"content" validate:"required"`
	Digest           string `json:"digest,omitempty"`
}

type MpNews struct {
	Articles []MpArticle `json:"articles" validate:"required"`
}

type MpNewsMessage struct {
	Message
	Safe   int    `json:"safe,omitempty" validate:"omitempty,oneof=0 1 2"`
	MpNews MpNews `json:"mpnews" validate:"required"`
}

type MarkDownMessage struct {
	Message
	MarkDown Text `json:"markdown" validate:"required"`
}

type MessageSendResponse struct {
	internal.BizResponse
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
	MsgId        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

func (app workChat) MessageSend(msg interface{}) (resp MessageSendResponse) {
	if ok := validate.Struct(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	h := H{}
	buf, _ := json.Marshal(msg)
	json.Unmarshal(buf, &h)
	h["agentid"] = app.appId
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
	}

	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/message/send?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}