package workchatapp

//参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92571

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
