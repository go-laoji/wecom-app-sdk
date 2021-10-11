package workchatapp

//
// 客户标签管理
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92116
//

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

type CorpTagGroup struct {
	GroupId    string    `json:"group_id"`
	GroupName  string    `json:"group_name"`
	CreateTime uint64    `json:"create_time,omitempty"`
	Order      int       `json:"order,omitempty"`
	Deleted    bool      `json:"deleted,omitempty"`
	Tag        []CorpTag `json:"tag"`
}

type CorpTag struct {
	Id         string `json:"id,omitempty"`
	Name       string `json:"name" validate:"required,max=30"`
	Order      int32  `json:"order"`
	CreateTime uint64 `json:"create_time,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
}

type CorpTagListResponse struct {
	internal.BizResponse
	TagGroup []CorpTagGroup `json:"tag_group"`
}

// CorpTagList 若tag_id和group_id均为空，则返回所有标签。
// 同时传递tag_id和group_id时，忽略tag_id，仅以group_id作为过滤条件。
func (app workChat) CorpTagList(tagIds, groupIds []string) (resp CorpTagListResponse) {
	p := H{"tag_id": tagIds, "group_id": groupIds}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_corp_tag_list?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type CorpTagAddResponse struct {
	internal.BizResponse
	TagGroup CorpTagGroup `json:"tag_group"`
}

// CorpTagAdd 企业可通过此接口向客户标签库中添加新的标签组和标签，每个企业最多可配置3000个企业标签。
func (app workChat) CorpTagAdd(tagGroup CorpTagGroup) (resp CorpTagAddResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/add_corp_tag?%s", queryParams.Encode()), tagGroup)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// CorpTagUpdate 企业可通过此接口编辑客户标签/标签组的名称或次序值。
func (app workChat) CorpTagUpdate(tag CorpTag) (resp internal.BizResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/edit_corp_tag?%s", queryParams.Encode()), tag)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// CorpTagDelete 企业可通过此接口删除客户标签库中的标签，或删除整个标签组。
func (app workChat) CorpTagDelete(tagIds, groupIds []string) (resp internal.BizResponse) {
	p := H{"tag_id": tagIds, "group_id": groupIds}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/del_corp_tag?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
