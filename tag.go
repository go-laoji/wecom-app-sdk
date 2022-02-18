package wecom

//
// 企业微信标签管理接口
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/90209

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-app-sdk/internal"
)

type Tag struct {
	TagId   int    `json:"tagid"`
	TagName string `json:"tagname" validate:"required,max=32"`
}

type TagCreateResponse struct {
	internal.BizResponse
	TagId int `json:"tagid"`
}

// TagCreate 创建标签
func (app weCom) TagCreate(tag Tag) (resp TagCreateResponse) {
	if ok := validate.Struct(tag); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/tag/create?%s", queryParams.Encode()), tag)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// TagUpdate 更新标签名字
func (app weCom) TagUpdate(tag Tag) (resp internal.BizResponse) {
	if ok := validate.Struct(tag); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/tag/update?%s", queryParams.Encode()), tag)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// TagDelete 删除标签
func (app weCom) TagDelete(id int) (resp internal.BizResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("tagid", fmt.Sprintf("%v", id))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/tag/delete?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type TagListResponse struct {
	internal.BizResponse
	TagList []Tag `json:"taglist"`
}

// TagList 获取标签列表
func (app weCom) TagList() (resp TagListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/tag/list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type TagUserListResponse struct {
	internal.BizResponse
	TagName  string `json:"tagname"`
	UserList []struct {
		UserId string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
	PartyList []int32 `json:"partylist"`
}

// TagUserList 获取标签成员
func (app weCom) TagUserList(id int) (resp TagUserListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("tagid", fmt.Sprintf("%v", id))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/tag/get?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type TagAddOrDelUsersResponse struct {
	internal.BizResponse
	InvalidList  string  `json:"invalidlist,omitempty"`
	InvalidParty []int32 `json:"invalidparty,omitempty"`
}

// TagAddUsers 增加标签成员
func (app weCom) TagAddUsers(tagId int, userIds []string, partyIds []int32) (resp TagAddOrDelUsersResponse) {
	p := H{"tagid": tagId, "userlist": userIds, "partylist": partyIds}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/tag/addtagusers?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// TagDelUsers 删除标签成员
func (app weCom) TagDelUsers(tagId int, userIds []string, partyIds []int32) (resp TagAddOrDelUsersResponse) {
	p := H{"tagid": tagId, "userlist": userIds, "partylist": partyIds}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/tag/deltagusers?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
