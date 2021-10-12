package workchatapp

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

type TransferCustomerRequest struct {
	HandoverUserId     string   `json:"handover_userid" validate:"required"`
	TakeoverUserId     string   `json:"takeover_userid" validate:"required"`
	ExternalUserId     []string `json:"external_userid" validate:"required"`
	TransferSuccessMsg string   `json:"transfer_success_msg" validate:"omitempty,max=200"`
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
func (app workChat) TransferCustomer(request TransferCustomerRequest) (resp TransferCustomerResponse) {
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
func (app workChat) TransferResult(request TransferResultRequest) (resp TransferResultResponse) {
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
