package wecom

//
// 客户联系规则组管理
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/94883
//
import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-app-sdk/internal"
)

type ExternalContactCustomerStrategyListResponse struct {
	internal.BizResponse
	Strategy []struct {
		StrategyId int `json:"strategy_id"`
	} `json:"strategy"`
	NextCursor string `json:"next_cursor"`
}

// ExternalContactCustomerStrategyList 获取规则组列表
// 企业可通过此接口获取企业配置的所有客户规则组id列表
func (app weCom) ExternalContactCustomerStrategyList(cursor string, limit int) (resp ExternalContactCustomerStrategyListResponse) {
	p := H{"cursor": cursor, "limit": limit}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/customer_strategy/list?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type Strategy struct {
	StrategyID   int      `json:"strategy_id,omitempty"`
	ParentID     int      `json:"parent_id"`
	StrategyName string   `json:"strategy_name" validate:"required"`
	CreateTime   int      `json:"create_time,omitempty"`
	AdminList    []string `json:"admin_list" validate:"required,max=20"`
	Privilege    struct {
		ViewCustomerList        bool `json:"view_customer_list"`
		ViewCustomerData        bool `json:"view_customer_data"`
		ViewRoomList            bool `json:"view_room_list"`
		ContactMe               bool `json:"contact_me"`
		JoinRoom                bool `json:"join_room"`
		ShareCustomer           bool `json:"share_customer"`
		OperResignCustomer      bool `json:"oper_resign_customer"`
		OperResignGroup         bool `json:"oper_resign_group"`
		SendCustomerMsg         bool `json:"send_customer_msg"`
		EditWelcomeMsg          bool `json:"edit_welcome_msg"`
		ViewBehaviorData        bool `json:"view_behavior_data"`
		ViewRoomData            bool `json:"view_room_data"`
		SendGroupMsg            bool `json:"send_group_msg"`
		RoomDeduplication       bool `json:"room_deduplication"`
		RapidReply              bool `json:"rapid_reply"`
		OnjobCustomerTransfer   bool `json:"onjob_customer_transfer"`
		EditAntiSpamRule        bool `json:"edit_anti_spam_rule"`
		ExportCustomerList      bool `json:"export_customer_list"`
		ExportCustomerData      bool `json:"export_customer_data"`
		ExportCustomerGroupList bool `json:"export_customer_group_list"`
		ManageCustomerTag       bool `json:"manage_customer_tag"`
	} `json:"privilege"`
}

type StrategyRange struct {
	Type    int    `json:"type" validate:"required,oneof=1 2"`
	UserId  string `json:"userid,omitempty"`
	PartyId int32  `json:"partyid,omitempty"`
}

type ExternalContactCustomerStrategyGetResponse struct {
	internal.BizResponse
	Strategy Strategy `json:"strategy"`
}

// ExternalContactCustomerStrategyGet 获取规则组详情
// 企业可以通过此接口获取某个客户规则组的详细信息。
func (app weCom) ExternalContactCustomerStrategyGet(strategyId int) (resp ExternalContactCustomerStrategyGetResponse) {
	p := H{"strategy_id": strategyId}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/customer_strategy/get?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ExternalContactCustomerStrategyGetRangeResponse struct {
	internal.BizResponse
	Range []struct {
		Type    int    `json:"type"`
		UserId  string `json:"userid,omitempty"`
		PartyId int32  `json:"partyid,omitempty"`
	} `json:"range"`
	NextCursor string `json:"next_cursor"`
}

// ExternalContactCustomerStrategyGetRange 获取规则组管理范围
func (app weCom) ExternalContactCustomerStrategyGetRange(strategyId int, cursor string, limit int) (resp ExternalContactCustomerStrategyGetRangeResponse) {
	p := H{"strategy_id": strategyId, "cursor": cursor, "limit": limit}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/customer_strategy/get_range?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ExternalContactCustomerStrategyCreateRequest struct {
	Strategy
	Range []StrategyRange `json:"range" validate:"required"`
}

type ExternalContactCustomerStrategyCreateResponse struct {
	internal.BizResponse
	StrategyId int `json:"strategy_id"`
}

// ExternalContactCustomerStrategyCreate 创建新的规则组
// 企业可通过此接口创建一个新的客户规则组。该接口仅支持串行调用，请勿并发创建规则组。
func (app weCom) ExternalContactCustomerStrategyCreate(request ExternalContactCustomerStrategyCreateRequest) (resp ExternalContactCustomerStrategyCreateResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/customer_strategy/create?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ExternalContactCustomerStrategyEditRequest struct {
	Strategy
	RangeAdd []StrategyRange `json:"range_add"`
	RangeDel []StrategyRange `json:"range_del"`
}

// ExternalContactCustomerStrategyEdit 编辑规则组及其管理范围
func (app weCom) ExternalContactCustomerStrategyEdit(request ExternalContactCustomerStrategyEditRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/customer_strategy/create?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// ExternalContactCustomerStrategyDelete 删除规则组
func (app weCom) ExternalContactCustomerStrategyDelete(strategyId int) (resp internal.BizResponse) {
	p := H{"strategy_id": strategyId}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/customer_strategy/del?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
