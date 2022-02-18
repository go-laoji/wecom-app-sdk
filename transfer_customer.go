package wecom

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-app-sdk/internal"
)

type TransferCustomerRequest struct {
	HandoverUserId     string   `json:"handover_userid" validate:"required"`
	TakeoverUserId     string   `json:"takeover_userid" validate:"required"`
	ExternalUserId     []string `json:"external_userid" validate:"required"`
	TransferSuccessMsg string   `json:"transfer_success_msg,omitempty" validate:"omitempty,max=200"`
}

type TransferCustomerResponse struct {
	internal.BizResponse
	Customer []struct {
		ExternalUserId string `json:"external_userid"`
		ErrCode        int    `json:"errcode"`
	}
}

// TransferCustomer 分配在职成员的客户
// 企业可通过此接口，转接在职成员的客户给其他成员。
// 参考连接 https://open.work.weixin.qq.com/api/doc/90000/90135/92125
func (app weCom) TransferCustomer(request TransferCustomerRequest) (resp TransferCustomerResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/transfer_customer?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type TransferResultRequest struct {
	HandoverUserId string `json:"handover_userid" validate:"required"`
	TakeoverUserId string `json:"takeover_userid" validate:"required"`
	Cursor         string `json:"cursor"`
}

type TransferResultResponse struct {
	internal.BizResponse
	Customer []struct {
		ExternalUserId string `json:"external_userid"`
		Status         int    `json:"status"`
		TakeoverTime   uint64 `json:"takeover_time"`
	} `json:"customer"`
	NextCursor string `json:"next_cursor"`
}

// TransferResult 查询客户接替状态
// 企业和第三方可通过此接口查询在职成员的客户转接情况。
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/94088
func (app weCom) TransferResult(request TransferResultRequest) (resp TransferResultResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/transfer_result?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UnAssignedRequest struct {
	PageId   int    `json:"page_id" validate:"required_without=Cursor,omitempty"`
	PageSize int    `json:"page_size" validate:"max=1000"`
	Cursor   string `json:"cursor" validate:"required_without=PageId,omitempty"`
}

type UnAssignedInfo struct {
	HandoverUserId string `json:"handover_userid"`
	ExternalUserId string `json:"external_userid"`
	DimissionTime  uint64 `json:"dimission_time"`
}

type UnAssignedResponse struct {
	internal.BizResponse
	Info       []UnAssignedInfo `json:"info"`
	IsLast     bool             `json:"is_last"`
	NextCursor string           `json:"next_cursor"`
}

// GetUnassignedList 获取待分配的离职成员列表
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92124
func (app weCom) GetUnassignedList(request UnAssignedRequest) (resp UnAssignedResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_unassigned_list?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// TransferCustomerResigned 分配离职成员的客户;不可设置　TransferSuccessMsg
// handover_userid必须是已离职用户
// external_userid必须是handover_userid的客户
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/94081
func (app weCom) TransferCustomerResigned(request TransferCustomerRequest) (resp TransferCustomerResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/resigned/transfer_customer?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// TransferResultResigned 查询客户接替状态
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/94082
func (app weCom) TransferResultResigned(request TransferResultRequest) (resp TransferResultResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/transfer_result?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GroupChatTransferRequest struct {
	ChatIdList []string `json:"chat_id_list"`
	NewOwner   string   `json:"new_owner"`
}

type GroupChatTransferResponse struct {
	internal.BizResponse
	FailedChatList []struct {
		ChatId  string `json:"chat_id"`
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	} `json:"failed_chat_list"`
}

// TransferGroupChat 分配离职成员的客户群
// 可通过此接口，将已离职成员为群主的群，分配给另一个客服成员
// 群主离职了的客户群，才可继承
// 继承给的新群主，必须是配置了客户联系功能的成员
// 继承给的新群主，必须有设置实名
// 继承给的新群主，必须有激活企业微信
// 同一个人的群，限制每天最多分配300个给新群主
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/92127
func (app weCom) TransferGroupChat(request GroupChatTransferRequest) (resp GroupChatTransferResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/groupchat/transfer?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
